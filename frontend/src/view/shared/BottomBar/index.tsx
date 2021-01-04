import React from 'react'
import LogoIconImage from '@assets/images/logo-icon.png'
import {
  BottomBarWrapper,
  CopyrightWrapper,
  FooterWrapper,
  LogoIcon,
  FooterMenu,
  FooterMenuItem,
  FooterMenuLink,
  FooterMenuALink,
  LogoWrapper,
} from './styles'

export const BottomBar = React.memo(() => {
  return (
    <BottomBarWrapper>
      <FooterWrapper>
        <LogoWrapper to="/messages">
          <LogoIcon src={LogoIconImage} />
        </LogoWrapper>
        <FooterMenu>
          <FooterMenuItem>
            <FooterMenuALink href="https://docs.koto.at/#/help">CONTACT US</FooterMenuALink>
          </FooterMenuItem>
          <FooterMenuItem>
            <FooterMenuALink href="https://docs.koto.at">ABOUT KOTO</FooterMenuALink>
          </FooterMenuItem>
          <FooterMenuItem>
            <FooterMenuLink to="/docs/code-of-conduct">TERMS OF USE</FooterMenuLink>
          </FooterMenuItem>
        </FooterMenu>
      </FooterWrapper>
      <CopyrightWrapper>Koto © Copyright 2021 | All Rights Reserved</CopyrightWrapper>
    </BottomBarWrapper>
  )
})