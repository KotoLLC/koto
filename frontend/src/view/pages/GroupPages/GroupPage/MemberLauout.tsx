import React, { useEffect, useState } from 'react'
import { connect } from 'react-redux'
import Actions from '@store/actions'
import selectors from '@selectors/index'
import { ApiTypes, StoreTypes } from 'src/types'
import { PageLayout } from '@view/shared/PageLayout'
import { GroupTopBar } from './GroupTopBar'
import { Member } from './Member'
import { Owner } from './Owner'
import AvatarIcon from '@assets/images/groups-avatar-icon.svg'
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
  ViewMoreButton,
} from './styles'

interface Props {
  groupDetails?: ApiTypes.Groups.GroupDetails | null
}

const MemberLauout: React.FC<Props> = React.memo((props) => {
  // const [isRequested, setRequested] = useState(false)
  const { groupDetails } = props

  // console.log(groupDetails)

  if (!groupDetails) return null

  const { group, members } = groupDetails

  return (
    <PageLayout>
      <GroupCover />
      <GroupTopBar groupId={''} isAdminLayout={false}/>
      <GroupContainer>
        <GroupMainWrapper>
          <LeftSideBar>
            <AvatarStyled>
              <img src={AvatarIcon} alt="icon" />
            </AvatarStyled>
            <GroupName>{group?.name}</GroupName>
            <GroupPublicity>{group?.is_public ? 'Public' : 'Private'} group</GroupPublicity>
            <GroupDescriptopn>{group?.description}</GroupDescriptopn>
          </LeftSideBar>
          <CentralBar>
            <BarTitle>Members ({members?.length})</BarTitle>
            {Boolean(members?.length) && members.map(item => (
              <Member
                groupId={group?.id}
                isAdminLayout={false}
                key={uuidv4()}
                {...item}
              />
            ))}
            {/* <ViewMoreButton>View more</ViewMoreButton> */}
          </CentralBar>
          <RightSideBar>
            <BarTitle>Owner</BarTitle>
            <Owner {...group.admin}/>
          </RightSideBar>
        </GroupMainWrapper>
      </GroupContainer>
    </PageLayout>
  )
})

type StateProps = Pick<Props, 'groupDetails' >
const mapStateToProps = (state: StoreTypes): StateProps => ({
  groupDetails: selectors.groups.groupDetails(state),
})

// type DispatchProps = Pick<Props, 'onGetInvitesToConfirmRequest'>
// const mapDispatchToProps = (dispatch): DispatchProps => ({
//   onGetInvitesToConfirmRequest: () => dispatch(Actions.groups.getInvitesToConfirmRequest()),
// })

export default connect(mapStateToProps, null)(MemberLauout)