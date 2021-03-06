import React from 'react'
import { connect } from 'react-redux'
import { StoreTypes } from 'src/types'
import selectors from '@selectors/index'
import {
  SidebarWrapper,
  SidebarItem,
  SidebarButtonWrapper,
  SidebarButtonLink,
} from '@view/shared/styles'

interface Props {
  ownedHubs: string[]
}

const GroupsSidebar: React.FC<Props> = React.memo((props) => {
  // const { ownedHubs } = props

  return (
    <SidebarWrapper>
      <SidebarItem exact to="/groups">PUBLIC GROUPS</SidebarItem>
      <SidebarItem to="/groups/my">MY GROUPS</SidebarItem>
      <SidebarButtonWrapper>
        <SidebarButtonLink to="/groups/create">Create New Group</SidebarButtonLink>
      </SidebarButtonWrapper>
    </SidebarWrapper>
  )
})

type StateProps = Pick<Props, 'ownedHubs'>
const mapStateToProps = (state: StoreTypes): StateProps => ({
  ownedHubs: selectors.profile.ownedHubs(state),
})

export default connect(mapStateToProps)(GroupsSidebar)