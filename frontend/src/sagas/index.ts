import { all, takeEvery } from 'redux-saga/effects'
import { Types as AuthorizationTypes } from '@store/authorization/actions'
import { Types as FriendTypes } from '@store/friends/actions'
import { Types as NodeTypes } from '@store/nodes/actions'
import { Types as ProfileTypes } from '@store/profile/actions'
import { Types as MessagesTypes } from '@store/messages/actions'

import { 
    watchlogin, 
    watchlogout,
    watchGetAuthToken,
} from './authorization'
import { 
    watchGetFriends, 
    watchGetFriendsOfFriends,
    watchAddFriend,
    watchGetInvitations,
    watchAcceptInvitation,
    watchRejectInvitation,
} from './friends'
import { 
    watchNodeCreate, 
    watchGetNodes, 
    watchApproveNode,
    watchRemoveNode,
} from './nodes'
import { watchGetProfile } from './profile'
import { 
    watchGetMessages, 
    watchGetNodeToken, 
} from './messages'

export function* rootSaga() {
    yield all([
        takeEvery(AuthorizationTypes.LOGIN_REQUEST, watchlogin),
        takeEvery(AuthorizationTypes.LOGOUT_REQUEST, watchlogout),
        takeEvery(AuthorizationTypes.GET_AUTH_TOKEN_REQUEST, watchGetAuthToken),
        takeEvery(FriendTypes.GET_FRIENDS_REQUEST, watchGetFriends),
        takeEvery(FriendTypes.GET_FRIENDS_OF_FRIENDS_REQUEST, watchGetFriendsOfFriends),
        takeEvery(FriendTypes.ADD_FRIEND_REQUEST, watchAddFriend),
        takeEvery(FriendTypes.GET_INVITATIONS_REQUEST, watchGetInvitations),
        takeEvery(FriendTypes.ACCEPT_INVITATION_REQUEST, watchAcceptInvitation),
        takeEvery(FriendTypes.REJECT_INVITATION_REQUEST, watchRejectInvitation),
        takeEvery(NodeTypes.NODE_CREATE_REQUEST, watchNodeCreate),
        takeEvery(NodeTypes.GET_NODES_REQUEST, watchGetNodes),
        takeEvery(NodeTypes.APPROVE_NODE_REQUEST, watchApproveNode),
        takeEvery(NodeTypes.REMOVE_NODE_REQUEST, watchRemoveNode),
        takeEvery(ProfileTypes.GET_PROFILE_REQUEST, watchGetProfile),
        takeEvery(MessagesTypes.GET_MESSAGES_REQUEST, watchGetMessages),
        takeEvery(MessagesTypes.GET_CURRENT_TOKEN_REQUEST, watchGetNodeToken),
    ])
}