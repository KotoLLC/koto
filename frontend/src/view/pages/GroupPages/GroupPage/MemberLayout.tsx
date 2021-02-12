import React from 'react'
import { connect } from 'react-redux'
import selectors from '@selectors/index'
import { ApiTypes, StoreTypes } from 'src/types'
import GroupTopBar from './GroupTopBar'
import { Member } from './Member'
import { Owner } from './Owner'
import { v4 as uuidv4 } from 'uuid'
import { getGroupAvatarUrl, getGroupCoverUrl } from '@services/avatarUrl'
import { Container } from '@view/shared/styles'
import {
  GroupCover,
  GroupMainWrapper,
  LeftSideBar,
  RightSideBar,
  CentralBar,
  AvatarStyled,
  GroupName,
  GroupPublicity,
  GroupDescriptopn,
  BarTitle,
} from './styles'

interface Props {
  groupDetails?: ApiTypes.Groups.GroupDetails | null
}

const MemberLayout: React.FC<Props> = React.memo((props) => {
  const { groupDetails } = props

  if (!groupDetails) return null

  const { group, members, status } = groupDetails

  return (
    <>
      <GroupCover resource={getGroupCoverUrl(group?.id)} />
      <GroupTopBar
        className="desktop-only"
        groupId={group?.id}
        isAdminLayout={false}
        memberStatus={status}
        isPublic={group?.is_public}
      />
      <Container>
        <GroupMainWrapper>
          <LeftSideBar>
            <AvatarStyled src={getGroupAvatarUrl(group?.id)} />
            <GroupName>{group?.name}</GroupName>
            <GroupPublicity>{group?.is_public ? 'Public' : 'Private'} group</GroupPublicity>
            <GroupDescriptopn>{group?.description}</GroupDescriptopn>
            <GroupTopBar
              className="mobile-only"
              groupId={group?.id}
              isPublic={group?.is_public}
              isAdminLayout={false}
              memberStatus={status}
            />
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
            <Owner {...group.admin} />
          </RightSideBar>
        </GroupMainWrapper>
      </Container>
    </>
  )
})

type StateProps = Pick<Props, 'groupDetails'>
const mapStateToProps = (state: StoreTypes): StateProps => ({
  groupDetails: selectors.groups.groupDetails(state),
})

export default connect(mapStateToProps, null)(MemberLayout)