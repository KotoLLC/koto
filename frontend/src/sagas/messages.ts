import { put, all, call, select } from 'redux-saga/effects'
import Actions from '@store/actions'
import { API } from '@services/api'
import { ApiTypes } from 'src/types'
import { currentNodeBack2Front } from '@services/dataTransforms/currentNodeTransform'
import { nodesForMessagesBack2Front } from '@services/dataTransforms/nodesForMessagesTransform'
import { Types as MessagesTypes } from '@store/messages/actions'
import selectors from '@selectors/index'

export function* watchGetMessages() {
  const response = yield API.messages.getMessages()

  if (response.status === 200) {
    const messageTokens = nodesForMessagesBack2Front(response.data?.tokens)
    yield put(Actions.messages.getMessagesSucces(messageTokens))

    const state = yield select()
    const nodesWithMessages = selectors.messages.nodesWithMessages(state)

    yield all(messageTokens.map(item => {

      let count
 
      if (nodesWithMessages.get(item.host)) {
        const currentNode = nodesWithMessages.get(item.host)
        if (currentNode?.messages?.length) {
          count = currentNode.messages.length
        }
      }

      return call(watchGetMessagesFromNode, {
        type: MessagesTypes.GET_MESSAGES_FROM_NODE_REQUEST,
        payload: {
          host: item.host,
          body: {
            token: item.token,
            count,
          }
        },
      })
    }

    ))
  }
}
export function* watchGetMoreMessages() {
  const response = yield API.messages.getMessages()

  if (response.status === 200) {
    const messageTokens = nodesForMessagesBack2Front(response.data?.tokens)
    yield put(Actions.messages.getMoreMessagesSucces(messageTokens))

    yield all(messageTokens.map(item => call(watchGetMoreMessagesFromNode, {
      type: MessagesTypes.GET_MORE_MESSAGES_FROM_NODE_REQUEST,
      payload: {
        host: item.host,
        body: {
          token: item.token,
        }
      },
    })))
  }
}

export function* watchGetCurrentNode() {
  const response = yield API.messages.getCurrentNode()

  if (response.status === 200) {
    yield put(Actions.messages.getCurrentNodeSucces(currentNodeBack2Front(response.data?.tokens)))
  }
}

export function* watchGetMoreMessagesFromNode(action: { type: string, payload: ApiTypes.Messages.MessagesFromNode }) {
  const state = yield select()
  const nodesWithMessages = selectors.messages.nodesWithMessages(state)

  if (nodesWithMessages.get(action.payload.host)) {
    const currentNode = nodesWithMessages.get(action.payload.host)
    action.payload.body.from = currentNode?.lastMessageDate || ''
    action.payload.body.count = '10'
  }

  const response = yield API.messages.getMessagesFromNode(action.payload)

  if (response.status === 200) {
    let resultData = []
    if (response.data?.messages?.length) {
      resultData = response.data?.messages.map(item => {
        item.sourceHost = action.payload.host
        item.messageToken = action.payload.body.token
        return item
      })
    }

    yield put(Actions.messages.getMessagesFromNodeSucces({
      node: action.payload.host,
      messages: resultData
    }))
  } else {
    if (response.error.response.status === 400) {
      yield put(Actions.authorization.getAuthTokenRequest())
      yield put(Actions.messages.getMessagesRequest())
    }
  }
}

export function* watchGetMessagesFromNode(action: { type: string, payload: ApiTypes.Messages.MessagesFromNode }) {
  const response = yield API.messages.getMessagesFromNode(action.payload)

  if (response.status === 200) {
    let resultData = []
    if (response.data?.messages?.length) {
      resultData = response.data?.messages.map(item => {
        item.sourceHost = action.payload.host
        item.messageToken = action.payload.body.token
        return item
      })
    }

    yield put(Actions.messages.getMessagesFromNodeSucces({
      node: action.payload.host,
      messages: resultData
    }))
  } else {
    if (response.error.response.status === 400) {
      yield put(Actions.authorization.getAuthTokenRequest())
      yield put(Actions.messages.getMessagesRequest())
    }
  }
}

export function* watchPostMessage(action: { type: string, payload: ApiTypes.Messages.PostMessage }) {
  const response = yield API.messages.postMessage(action.payload)

  if (response.status === 200) {
    yield put(Actions.messages.postMessageSucces(true))
    yield put(Actions.messages.getMessagesRequest())
  }
}

export function* watchDeleteMessage(action: { type: string, payload: ApiTypes.Messages.DeleteMessage }) {
  const response = yield API.messages.deleteMessage(action.payload)

  if (response.status === 200) {
    yield put(Actions.messages.deleteMessageSucces())
    yield put(Actions.messages.getMessagesRequest())
  } else {
    yield put(Actions.common.setErrorNotify(response?.error?.response?.data?.msg || 'Server error'))
  }
}

export function* watchEditMessage(action: { type: string, payload: ApiTypes.Messages.EditMessage }) {
  const response = yield API.messages.editMessage(action.payload)

  if (response.status === 200) {
    yield put(Actions.messages.editMessageSucces())
    yield put(Actions.messages.getMessagesRequest())
  } else {
    yield put(Actions.common.setErrorNotify(response?.error?.response?.data?.msg || 'Server error'))
  }
}

export function* watchPostComment(action: { type: string, payload: ApiTypes.Messages.PostComment }) {
  const response = yield API.messages.postComment(action.payload)

  if (response.status === 200) {
    yield put(Actions.messages.postCommentSucces(true))
    yield put(Actions.messages.getMessagesRequest())
  } else {
    yield put(Actions.common.setErrorNotify(response?.error?.response?.data?.msg || 'Server error'))
  }
}

export function* watchEditComment(action: { type: string, payload: ApiTypes.Messages.EditComment }) {
  const response = yield API.messages.editComment(action.payload)

  if (response.status === 200) {
    yield put(Actions.messages.editCommentSucces())
    yield put(Actions.messages.getMessagesRequest())
  } else {
    yield put(Actions.common.setErrorNotify(response?.error?.response?.data?.msg || 'Server error'))
  }
}

export function* watchDeleteComment(action: { type: string, payload: ApiTypes.Messages.DeleteComment }) {
  const response = yield API.messages.deleteComment(action.payload)

  if (response.status === 200) {
    yield put(Actions.messages.deleteCommentSucces())
    yield put(Actions.messages.getMessagesRequest())
  } else {
    yield put(Actions.common.setErrorNotify(response?.error?.response?.data?.msg || 'Server error'))
  }
}

export function* watchGetMessageUploadLink(action: { type: string, payload: ApiTypes.Messages.UploadLinkRequest }) {
  const response = yield API.messages.getUploadLink(action.payload)

  if (response.status === 200) {
    yield put(Actions.messages.getMessageUploadLinkSucces(response.data))
  } else {
    yield put(Actions.common.setErrorNotify(response?.error?.response?.data?.msg || 'Server error'))
  }
}

export function* watchSetAttachment(action: { type: string, payload: ApiTypes.Messages.Attachment }) {
  const response = yield API.messages.setAttachment(action.payload.link, action.payload.form_data)

  if (response.status === 204 || response.status === 200) {
    yield put(Actions.messages.setAttachmentSuccess())
  } else {
    yield put(Actions.common.setErrorNotify(response?.error?.response?.data?.msg || 'Server error'))
  }
}

export function* watchLikeMessage(action: { type: string, payload: ApiTypes.Messages.Like }) {
  const response = yield API.messages.likeMessage(action.payload)

  if (response.status === 200) {
    yield put(Actions.messages.linkMessageSuccess())
    yield put(Actions.messages.getMessagesRequest())
    yield put(Actions.messages.getLikesForMessageRequest(action.payload))
  } else {
    yield put(Actions.common.setErrorNotify(response?.error?.response?.data?.msg || 'Server error'))
  }
}

export function* watchLikeComment(action: { type: string, payload: ApiTypes.Messages.Like }) {
  const response = yield API.messages.likeComment(action.payload)

  if (response.status === 200) {
    yield put(Actions.messages.linkCommentSuccess())
    yield put(Actions.messages.getMessagesRequest())
  } else {
    yield put(Actions.common.setErrorNotify(response?.error?.response?.data?.msg || 'Server error'))
  }
}

export function* watchGetLikesForMessage(action: { type: string, payload: ApiTypes.Messages.Like }) {
  const response = yield API.messages.getlikesForMessage(action.payload)

  if (response.status === 200) {
    yield put(Actions.messages.getLikesForMessageSuccess({
      id: action.payload.id,
      likes: response.data?.likes || []
    }))
  } else {
    yield put(Actions.common.setErrorNotify(response?.error?.response?.data?.msg || 'Server error'))
  }
}

export function* watchGetLikesForComment(action: { type: string, payload: ApiTypes.Messages.Like }) {
  const response = yield API.messages.getlikesForComment(action.payload)

  if (response.status === 200) {
    yield put(Actions.messages.getLikesForCommentSuccess({
      id: action.payload.id,
      likes: response.data?.likes || []
    }))
  } else {
    yield put(Actions.common.setErrorNotify(response?.error?.response?.data?.msg || 'Server error'))
  }
}
