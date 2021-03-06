import React, { useState, useEffect } from 'react'
import IconButton from '@material-ui/core/IconButton'
import Tooltip from '@material-ui/core/Tooltip'
import moment from 'moment'
import { connect } from 'react-redux'
import Actions from '@store/actions'
import selectors from '@selectors/index'
import SendIcon from '@material-ui/icons/Send'
import { getAvatarUrl } from '@services/avatarUrl'
import { Link } from 'react-router-dom'
import { AuthorButtonsMenu } from './AuthorButtonsMenu'
import { NoAuthorButtonsMenu } from './NoAuthorButtonsMenu'
import { ApiTypes, StoreTypes } from 'src/types'
import { LinkRenderer } from '@view/shared/LinkRenderer'
import ReactMarkdown from 'react-markdown'
import { MentionsInput, Mention } from 'react-mentions'
import { friendsToMentionFriends, MentionFriend } from '@services/dataTransforms/friendsToMentionFriends'
import {
  CommentWrapper,
  UserNameLink,
  CommentReactionsNav,
  CommentTextWrapper,
  AvatarStyled,
  EditMessageField,
  CircularProgressStyled,
  CommentContent,
  CommentReactionsNavWrapper,
  LikeCommentButton,
} from './styles'

interface Props extends ApiTypes.Feed.Comment {
  userId: string
  currentCommentLikes: ApiTypes.Feed.LikesInfoData | null
  friends: ApiTypes.Friends.Friend[] | null

  onCommentEdit: (data: ApiTypes.Feed.EditComment) => void
  onCommentDelete: (data: ApiTypes.Feed.DeleteComment) => void
  onLikeComment: (data: ApiTypes.Feed.Like) => void
  getLikesForComment: (data: ApiTypes.Feed.Like) => void
}

const Comment: React.SFC<Props> = (props) => {
  const {
    text,
    user_name,
    user_full_name,
    created_at,
    id,
    user_id,
    sourceHost,
    userId,
    liked_by_me,
    likes,
    onLikeComment,
    currentCommentLikes,
    getLikesForComment,
    friends,
  } = props
  const [isEditer, setEditor] = useState<boolean>(false)
  const [comment, onCommentChange] = useState<string>(text)
  const commentRef = React.createRef<HTMLDivElement>()
  const [isLikesInfoRequested, setLikesInfoRequest] = useState<boolean>(false)
  const [mentionFriends, setMentionFriends] = useState<MentionFriend[]>([])

  useEffect(() => {
    if (props.currentCommentLikes?.id === id) {
      setLikesInfoRequest(false)
    }
    if (!mentionFriends?.length && friends?.length) {
      setMentionFriends(friendsToMentionFriends(friends))
    }
  }, [props.currentCommentLikes, id, friends])

  const onMessageSave = () => {
    setEditor(false)
    props.onCommentEdit({
      host: sourceHost,
      body: {
        comment_id: id,
        text: comment,
        text_changed: true,
      }
    })
  }

  const onComandEnterDown = (event) => {
    if (event.keyCode === 13 && (event.metaKey || event.ctrlKey)) {
      onMessageSave()
    }
  }

  const getLikesInfo = () => {
    if (currentCommentLikes?.id === id) {
      setLikesInfoRequest(false)
    }

    if (currentCommentLikes?.id !== id) {
      setLikesInfoRequest(true)
      getLikesForComment({
        host: sourceHost,
        id: id
      })
    }
  }

  const rendreLikeButton = () => {
    let likesInfo = 'No likes yet'
    let usersLikes = ''

    if (currentCommentLikes?.id === id) {
      currentCommentLikes.likes.length && currentCommentLikes.likes.forEach((item, counter) => {

        if (counter < 15) {
          const likedUserName = item.user_full_name || item.user_name
          const comma = ((currentCommentLikes.likes.length - 1) === counter) ? '' : ', '
          usersLikes += `${likedUserName}${comma}`
        }

        if (counter === 15) {
          usersLikes += `...`
        }

      })
    }

    return (
      <Tooltip
        onClick={() => {
          if (liked_by_me) {
            onLikeComment({ host: sourceHost, id: id, unlike: true })
          } else {
            onLikeComment({ host: sourceHost, id: id })
          }
        }}
        title={(isLikesInfoRequested) ? <CircularProgressStyled size={30} /> : <>{usersLikes || likesInfo}</>}
        interactive onOpen={() => getLikesInfo()}>
        <LikeCommentButton>{likes} like</LikeCommentButton>
      </Tooltip>
    )

  }

  const renderCurrentIcons = () => {
    return (userId === user_id) ?
      <AuthorButtonsMenu {...{ message: comment, id, sourceHost, setEditor, isEditer }} /> :
      <NoAuthorButtonsMenu {...{ message: comment, id, sourceHost, }} />
  }

  return (
    <CommentWrapper ref={commentRef}>
      <Link to={`/profile/user?id=${user_id}`}>
        <AvatarStyled src={getAvatarUrl(user_id)} />
      </Link>
      <CommentTextWrapper>{
        isEditer ?
          <EditMessageField>
            <MentionsInput
              className="mentions"
              value={comment}
              onChange={(evant) => onCommentChange(evant.target.value)}
              onKeyDown={onComandEnterDown}
            >
              <Mention
                trigger="@"
                data={mentionFriends}
                className={'mentions__mention'}
                markup="[@__display__](/profile/user?id=__id__)"
              />
            </MentionsInput>
            <IconButton onClick={onMessageSave}>
              <SendIcon fontSize="small" />
            </IconButton>
          </EditMessageField>
          :
          <CommentContent className="markdown-body">
            <UserNameLink to={`/profile/user?id=${user_id}`}>{user_full_name || user_name}</UserNameLink>
            <ReactMarkdown renderers={{ link: LinkRenderer }}>{comment}</ReactMarkdown>
          </CommentContent>
      }
        <CommentReactionsNavWrapper>
          <CommentReactionsNav>{rendreLikeButton()}</CommentReactionsNav>
          <CommentReactionsNav>{moment(created_at).fromNow()}</CommentReactionsNav>
        </CommentReactionsNavWrapper>
      </CommentTextWrapper>
      {renderCurrentIcons()}
    </CommentWrapper>
  )
}

type StateProps = Pick<Props, 'userId' | 'currentCommentLikes' | 'friends'>
const mapStateToProps = (state: StoreTypes): StateProps => ({
  userId: selectors.profile.userId(state),
  currentCommentLikes: selectors.feed.currentCommentLikes(state),
  friends: selectors.friends.friends(state),
})

type DispatchProps = Pick<Props, 'onCommentEdit' | 'onCommentDelete' | 'onLikeComment' | 'getLikesForComment'>
const mapDispatchToProps = (dispatch): DispatchProps => ({
  onCommentEdit: (data: ApiTypes.Feed.EditComment) => dispatch(Actions.feed.editFeedCommentRequest(data)),
  onCommentDelete: (data: ApiTypes.Feed.DeleteComment) => dispatch(Actions.feed.deleteCommentRequest(data)),
  onLikeComment: (data: ApiTypes.Feed.Like) => dispatch(Actions.feed.linkFeedCommnetRequest(data)),
  getLikesForComment: (data: ApiTypes.Feed.Like) => dispatch(Actions.feed.getLikesForFeedCommentRequest(data)),
})

export default connect(mapStateToProps, mapDispatchToProps)(Comment)