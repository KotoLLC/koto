import React from 'react'
import {
  HubSettingsBlock,
  CreationHubWrapper,
  CreationHubTitle,
  CreationHubStepsWrapper,
  CreationHubStepWrapper,
  CreationHubStepIcon,
  CreationHubStepDescription,
  CreationHubNote,
} from './styles'
import { CommonTypes } from 'src/types'
import RemoveHubDialog from './RemoveHubDialog'
import StepIcon1 from '@assets/images/hub-step-icon-1.svg'
import StepIcon2 from '@assets/images/hub-step-icon-2.svg'
import StepIcon3 from '@assets/images/hub-step-icon-3.svg'

interface Props {
  isHubActive: boolean
  myActiveHub: CommonTypes.HubTypes.Hub
}

export const HubStepsInfo: React.FC<Props> = React.memo((props) => {
  const { isHubActive, myActiveHub } = props

  return (
    <HubSettingsBlock>
      <CreationHubWrapper>
        {!isHubActive && <CreationHubTitle>Create your own message hub!</CreationHubTitle>}
        <CreationHubStepsWrapper>
          <CreationHubStepWrapper>
            <CreationHubStepIcon src={StepIcon1} />
            <CreationHubStepDescription>Store your own messages, photos, and videos</CreationHubStepDescription>
          </CreationHubStepWrapper>
          <CreationHubStepWrapper>
            <CreationHubStepIcon src={StepIcon2} />
            <CreationHubStepDescription>Provide storage for your network of friends</CreationHubStepDescription>
          </CreationHubStepWrapper>
          <CreationHubStepWrapper>
            <CreationHubStepIcon src={StepIcon3} />
            <CreationHubStepDescription>Host your server on any cloud</CreationHubStepDescription>
          </CreationHubStepWrapper>
        </CreationHubStepsWrapper>
        {!isHubActive ?
          <CreationHubNote>
            Follow the instructions below to create a hub
        </CreationHubNote> :
          <RemoveHubDialog {...myActiveHub} />
        }
      </CreationHubWrapper>
    </HubSettingsBlock>
  )
})