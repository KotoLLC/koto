import React, { useEffect, useState } from 'react'
import { PageLayout } from '@view/shared/PageLayout'
import GroupTopBar from './GroupTopBar'
import { Member } from './Member'
import MemberInvited from './MemberInvited'
import AvatarIcon from '@assets/images/groups-avatar-icon.svg'
import DeleteGroupDialog from './DeleteGroupDialog'
import { connect } from 'react-redux'
import Actions from '@store/actions'
import selectors from '@selectors/index'
import { ApiTypes, StoreTypes } from 'src/types'
import { v4 as uuidv4 } from 'uuid'
import {
  GroupCover,
  GroupContainer,
  GroupMainWrapper,
  LeftSideBar,
  RightSideBar,
  CentralBar,
  AvatarStyled,
  GroupName,
  GroupPublicity,
  GroupDescriptopn,
  BarTitle,
  // ViewMoreButton,
} from './styles'

interface Props {
  groupDetails: ApiTypes.Groups.GroupDetails | null
  friends: ApiTypes.Friends.Friend[] | null

  onGetFriends: () => void
  onGetInvitesToConfirmRequest: () => void
}

const AdminLayoutPrivate: React.FC<Props> = React.memo((props) => {
  const [groupInvites, setGroupInvites] = useState<ApiTypes.Groups.Invite[] | null>(null)
  const [isRequested, setRequested] = useState(false)
  const { 
    groupDetails, 
    onGetInvitesToConfirmRequest, 
    onGetFriends,
    friends,
   } = props

  useEffect(() => {
    if (groupInvites === null && !isRequested) {
      onGetInvitesToConfirmRequest()
      setRequested(true)
    }

    if (groupDetails?.invites?.length && isRequested) {
      setGroupInvites(fixInvitesGroupId())
      setRequested(false)
    }

    if (friends === null) {
      onGetFriends()
    }

  }, [groupInvites, isRequested, friends])

  const fixInvitesGroupId = () => {
    if (!groupDetails?.invites?.length) return []
    
    return groupDetails?.invites?.map(item => {
      item.group_id = groupDetails?.group?.id
      return item
    })
  }

  if (!groupDetails) return null

  const { group, members, status, invites } = groupDetails

  return (
    <PageLayout>
      <GroupCover />
      <GroupTopBar 
        memberStatus={status}
        membersCounter={members?.length} 
        invitesCounter={invites?.length || 0} 
        groupId={group?.id} 
        isAdminLayout={true}
      />
      <GroupContainer>
        <GroupMainWrapper>
          <LeftSideBar>
            <AvatarStyled>
              <img src={AvatarIcon} alt="icon" />
            </AvatarStyled>
            <GroupName>{group?.name}</GroupName>
            <GroupPublicity>{group?.is_public ? 'Public' : 'Private'} group</GroupPublicity>
            <GroupDescriptopn>{group?.description}</GroupDescriptopn>
            <DeleteGroupDialog groupId={group?.id} />
          </LeftSideBar>
          <CentralBar>
            <BarTitle>Members ({members?.length})</BarTitle>
            {Boolean(members?.length) && members.map(item => (
              <Member
                groupId={group?.id}
                isAdminLayout={true}
                key={uuidv4()}
                {...item}
              />
            ))}
            {/* <ViewMoreButton>View more</ViewMoreButton> */}
          </CentralBar>
          <RightSideBar>
            <BarTitle>Invite friends</BarTitle>
            {Boolean(invites?.length) && invites?.map(item => <MemberInvited
              key={uuidv4()}
              {...item}
            />)}
            {/* <ViewMoreButton>View more</ViewMoreButton> */}
          </RightSideBar>
        </GroupMainWrapper>
      </GroupContainer>
    </PageLayout>
  )
})

type StateProps = Pick<Props, 'groupDetails' | 'friends'>
const mapStateToProps = (state: StoreTypes): StateProps => ({
  groupDetails: selectors.groups.groupDetails(state),
  friends: selectors.friends.friends(state),
})

type DispatchProps = Pick<Props, 'onGetInvitesToConfirmRequest' | 'onGetFriends'>
const mapDispatchToProps = (dispatch): DispatchProps => ({
  onGetInvitesToConfirmRequest: () => dispatch(Actions.groups.getInvitesToConfirmRequest()),
  onGetFriends: () => dispatch(Actions.friends.getFriendsRequest()),
})

export default connect(mapStateToProps, mapDispatchToProps)(AdminLayoutPrivate)
