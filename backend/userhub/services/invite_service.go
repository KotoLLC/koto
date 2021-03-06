package services

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"net/url"
	"strings"
	"time"

	"github.com/ansel1/merry"
	"github.com/twitchtv/twirp"

	"github.com/mreider/koto/backend/common"
	"github.com/mreider/koto/backend/userhub/repo"
	"github.com/mreider/koto/backend/userhub/rpc"
)

const (
	registerFrontendPath    = "/registration?email=%s&invite=%s"
	invitationsFrontendPath = "/friends/invitations"
)

type inviteService struct {
	*BaseService
}

func NewInvite(base *BaseService) rpc.InviteService {
	return &inviteService{
		BaseService: base,
	}
}

func (s *inviteService) Create(ctx context.Context, r *rpc.InviteCreateRequest) (*rpc.Empty, error) {
	me := s.getMe(ctx)
	meInfo := s.userCache.UserFullAccess(me.ID)
	if r.Friend == "" || r.Friend == me.ID || r.Friend == meInfo.Name || r.Friend == meInfo.Email {
		return nil, twirp.NewError(twirp.InvalidArgument, "")
	}

	friend := s.repos.User.FindUserByIDOrName(r.Friend)
	if friend == nil {
		if strings.Contains(r.Friend, "@") && strings.Contains(r.Friend, ".") {
			friends := s.repos.User.FindUsersByEmail(r.Friend)
			if len(friends) == 1 {
				friend = &friends[0]
			} else if len(friends) > 1 {
				return nil, twirp.NewError(twirp.AlreadyExists, "This email belongs to more than one account. Invite this person using their username instead.")
			}
		} else {
			return nil, twirp.NotFoundError("user not found")
		}
	}

	if friend != nil {
		areBlocked := s.repos.User.AreBlocked(me.ID, friend.ID)
		if areBlocked {
			return nil, twirp.NotFoundError("user not found")
		}

		if s.repos.Friend.AreFriends(me.ID, friend.ID) {
			return nil, twirp.NewError(twirp.AlreadyExists, "already a friend.")
		}

		s.repos.Invite.AddInvite(me.ID, friend.ID)
		if s.notificationSender != nil {
			s.notificationSender.SendNotification([]string{friend.ID}, meInfo.DisplayName+" invited you to be friends", "friend-invite/add", map[string]interface{}{
				"user_id": me.ID,
			})
		}
		friendInfo := s.userCache.UserFullAccess(friend.ID)
		err := s.sendInviteLinkToRegisteredUser(ctx, me, friendInfo.Email)
		if err != nil {
			log.Println("can't invite by email:", err)
		}
	} else {
		s.repos.Invite.AddInviteByEmail(me.ID, r.Friend)
		err := s.sendInviteLinkToUnregisteredUser(ctx, me, r.Friend)
		if err != nil {
			log.Println("can't invite by email:", err)
		}
	}

	return &rpc.Empty{}, nil
}

func (s *inviteService) Accept(ctx context.Context, r *rpc.InviteAcceptRequest) (*rpc.Empty, error) {
	me := s.getMe(ctx)
	if !s.repos.Invite.AcceptInvite(r.InviterId, me.ID, false) {
		return nil, twirp.NotFoundError("invite not found")
	}
	userInfo := s.userCache.User(me.ID, r.InviterId)
	if s.notificationSender != nil {
		s.notificationSender.SendNotification([]string{r.InviterId}, userInfo.DisplayName+" accepted your invitation!", "friend-invite/accept", map[string]interface{}{
			"user_id": me.ID,
		})
	}
	return &rpc.Empty{}, nil
}

func (s *inviteService) Reject(ctx context.Context, r *rpc.InviteRejectRequest) (*rpc.Empty, error) {
	me := s.getMe(ctx)
	if !s.repos.Invite.RejectInvite(r.InviterId, me.ID) {
		return nil, twirp.NotFoundError("invite not found")
	}
	if s.notificationSender != nil {
		s.notificationSender.SendNotification([]string{r.InviterId}, "Your invitation was rejected", "friend-invite/reject", map[string]interface{}{
			"user_id": me.ID,
		})
	}
	return &rpc.Empty{}, nil
}

func (s *inviteService) FromMe(ctx context.Context, _ *rpc.Empty) (*rpc.InviteFromMeResponse, error) {
	me := s.getMe(ctx)
	invites := s.repos.Invite.InvitesFromMe(me)
	rpcInvites := make([]*rpc.InviteFriendInvite, len(invites))
	for i, invite := range invites {
		invitedInfo := s.userCache.User(invite.FriendID, me.ID)
		friendName, friendFullName := invitedInfo.Name, invitedInfo.FullName
		if invite.FriendID == "" {
			friendName = invite.FriendEmail
			friendFullName = ""
		}

		rpcInvites[i] = &rpc.InviteFriendInvite{
			FriendId:       invite.FriendID,
			FriendName:     friendName,
			FriendFullName: friendFullName,
			CreatedAt:      common.TimeToRPCString(invite.CreatedAt),
			AcceptedAt:     common.NullTimeToRPCString(invite.AcceptedAt),
			RejectedAt:     common.NullTimeToRPCString(invite.RejectedAt),
		}
	}

	return &rpc.InviteFromMeResponse{
		Invites: rpcInvites,
	}, nil
}

func (s *inviteService) ForMe(ctx context.Context, _ *rpc.Empty) (*rpc.InviteForMeResponse, error) {
	me := s.getMe(ctx)
	invites := s.repos.Invite.OpenInvitesForMe(me)
	rpcInvites := make([]*rpc.InviteFriendInvite, len(invites))
	for i, invite := range invites {
		inviterInfo := s.userCache.UserFullAccess(invite.UserID)
		rpcInvites[i] = &rpc.InviteFriendInvite{
			FriendId:       invite.UserID,
			FriendName:     inviterInfo.Name,
			FriendFullName: inviterInfo.FullName,
			CreatedAt:      common.TimeToRPCString(invite.CreatedAt),
			AcceptedAt:     common.NullTimeToRPCString(invite.AcceptedAt),
			RejectedAt:     common.NullTimeToRPCString(invite.RejectedAt),
		}
	}

	return &rpc.InviteForMeResponse{
		Invites: rpcInvites,
	}, nil
}

func (s *inviteService) sendInviteLinkToUnregisteredUser(ctx context.Context, inviter repo.User, userEmail string) error {
	if !s.mailSender.Enabled() {
		return nil
	}

	inviteToken, err := s.tokenGenerator.Generate(inviter.ID, "user-invite",
		time.Now().Add(time.Hour*24*30*12),
		map[string]interface{}{
			"email": userEmail,
		})
	if err != nil {
		return merry.Wrap(err)
	}

	attachments := s.GetUserAttachments(ctx, inviter.ID)

	inviterInfo := s.userCache.UserFullAccess(inviter.ID)
	link := fmt.Sprintf("%s"+registerFrontendPath, s.cfg.FrontendAddress, url.QueryEscape(userEmail), inviteToken)
	var message bytes.Buffer
	err = s.rootEmailTemplate.ExecuteTemplate(&message, "friend_request.gohtml", map[string]interface{}{
		"InviterDisplayName": inviterInfo.DisplayName,
		"AcceptLink":         link,
		"RejectLink":         link,
	})
	if err != nil {
		return err
	}
	return s.mailSender.SendHTMLEmail([]string{userEmail}, inviterInfo.DisplayName+" invited you to be friends",
		message.String(), attachments)
}

func (s *inviteService) sendInviteLinkToRegisteredUser(ctx context.Context, inviter repo.User, userEmail string) error {
	if !s.mailSender.Enabled() {
		return nil
	}

	attachments := s.GetUserAttachments(ctx, inviter.ID)

	inviterInfo := s.userCache.UserFullAccess(inviter.ID)
	link := fmt.Sprintf("%s"+invitationsFrontendPath, s.cfg.FrontendAddress)
	var message bytes.Buffer
	err := s.rootEmailTemplate.ExecuteTemplate(&message, "friend_request.gohtml", map[string]interface{}{
		"InviterDisplayName": inviterInfo.DisplayName,
		"AcceptLink":         link,
		"RejectLink":         link,
	})
	if err != nil {
		return err
	}

	return s.mailSender.SendHTMLEmail([]string{userEmail}, inviterInfo.DisplayName+" invited you to be friends",
		message.String(), attachments)
}
