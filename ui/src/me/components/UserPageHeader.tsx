// Libraries
import React, {PureComponent} from 'react'
import {CLOUD} from 'src/shared/constants'

// Components
import {
  Page,
  FlexBox,
  FlexDirection,
  AlignItems,
  ComponentSize,
} from '@influxdata/clockface'

// Constants
import {generateRandomGreeting} from 'src/me/constants'

interface Props {
  userName: string
  orgName: string
}

export default class UserPageHeader extends PureComponent<Props> {
  public render() {
    return (
      <Page.ControlBar fullWidth={false}>
        <Page.ControlBarLeft>
          <FlexBox
            direction={FlexDirection.Column}
            alignItems={AlignItems.FlexStart}
            margin={ComponentSize.Small}
          >
            {this.title}
          </FlexBox>
        </Page.ControlBarLeft>
        <Page.ControlBarRight />
      </Page.ControlBar>
    )
  }

  private get title(): JSX.Element {
    const {userName, orgName} = this.props

    const {text, language} = generateRandomGreeting()

    let title = ''

    if (CLOUD) {
      title = `${text}, ${userName}! Welcome to InfluxDB Cloud!`
    } else {
      title = `${text}, ${userName}! Welcome to ${orgName}!`
    }

    const altText = `That's how you say hello in ${language}`

    return <Page.Title title={title} altText={altText} />
  }
}
