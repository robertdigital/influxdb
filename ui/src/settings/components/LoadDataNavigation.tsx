// Libraries
import React, {PureComponent} from 'react'
import _ from 'lodash'
import {withRouter, WithRouterProps} from 'react-router'

// Components
import {
  Tabs,
  Orientation,
  ComponentSize,
} from '@influxdata/clockface'
import {FeatureFlag} from 'src/shared/utils/featureFlag'

// Decorators
import {ErrorHandling} from 'src/shared/decorators/errors'
import CloudExclude from 'src/shared/components/cloud/CloudExclude'

interface OwnProps {
  activeTab: string
  orgID: string
}

type Props = OwnProps & WithRouterProps

@ErrorHandling
class LoadDataNavigation extends PureComponent<Props> {
  public render() {
    const {activeTab, orgID, router} = this.props

    const handleTabClick = (id: string): void => {
      router.push(`/orgs/${orgID}/load-data/${id}`)
    }

    const tabs = [
      {
        text: 'Buckets',
        id: 'buckets',
        cloudExclude: false,
        featureFlag: null,
      },
      {
        text: 'Telegraf',
        id: 'telegrafs',
        cloudExclude: false,
        featureFlag: null,
      },
      {
        text: 'Scrapers',
        id: 'scrapers',
        cloudExclude: true,
        featureFlag: null,
      },
      {
        text: 'Tokens',
        id: 'tokens',
        cloudExclude: false,
        featureFlag: null,
      },
      {
        text: 'Client Libraries',
        id: 'client-libraries',
        cloudExclude: false,
        featureFlag: null,
      },
    ]

    return (
      <Tabs
        orientation={Orientation.Horizontal}
        size={ComponentSize.Large}
      >
        {tabs.map(t => {
          let tabElement = (
            <Tabs.Tab
              key={t.id}
              text={t.text}
              id={t.id}
              onClick={handleTabClick}
              active={t.id === activeTab}
            />
          )

          if (t.cloudExclude) {
            tabElement = <CloudExclude key={t.id}>{tabElement}</CloudExclude>
          }

          if (t.featureFlag) {
            tabElement = (
              <FeatureFlag key={t.id} name={t.featureFlag}>
                {tabElement}
              </FeatureFlag>
            )
          }
          return tabElement
        })}
      </Tabs>
    )
  }
}

export default withRouter(LoadDataNavigation)
