import styled from 'styled-components'
import List from '@material-ui/core/List'
import SearchIcon from '@material-ui/icons/Search'
import Paper from '@material-ui/core/Paper'
import IconButton from '@material-ui/core/IconButton'

export const SidebarWrapper = styled.aside`
  max-width: 400px;
  width: 100%;
  margin-right: 20px;

  @media (max-width: 1024px) {
    max-width: 350px;
  }
`

export const ContentWrapper = styled(Paper)`
  width: calc(100% - 420px);
  padding: 10px 15px;

  @media (max-width: 1024px) {
    width: calc(100% - 370px);
  }
`

export const ListStyled = styled(List)`
  position: relative;
  overflow: auto;
  height: calc(100vh - 220px);
`

export const SearchWrapper = styled.div`
  background: #fff;
  width: 100%;
  padding: 10px 0px 0;
  border-radius: 4px 4px 0 0;
`

export const SearchIconStyled = styled(SearchIcon)`
  && {
    margin-left: 10px;
  }
`

export const ContainerTitle = styled.h3`
  font-size: 14px;
  padding: 0 0 5px;
  margin: 0;
  text-transform: uppercase;
`

export const EmptyMessage = styled.div`
  padding: 15px;
  width: 100%;
  display: flex;
  line-height: center;
  align-content: center;
`

export const UserNoteUnderlined = styled.span`
  text-decoration: underline;
  text-transform: capitalize;
  color: #1976d2;
  cursor: pointer;
`
export const UserName = styled.span`
  text-transform: capitalize;
`

export const IconButtonGreen = styled(IconButton)`
  && {
    color: #1fc456;
  }
`