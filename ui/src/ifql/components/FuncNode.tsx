import React, {PureComponent, MouseEvent} from 'react'

import FuncArgs from 'src/ifql/components/FuncArgs'
import FuncArgsPreview from 'src/ifql/components/FuncArgsPreview'
import {OnDeleteFuncNode, OnChangeArg, Func} from 'src/types/ifql'
import {ErrorHandling} from 'src/shared/decorators/errors'
import {Service} from 'src/types'

interface Props {
  func: Func
  service: Service
  bodyID: string
  declarationID?: string
  onDelete: OnDeleteFuncNode
  onChangeArg: OnChangeArg
  onGenerateScript: () => void
  declarationsFromBody: string[]
}

interface State {
  isExpanded: boolean
}

@ErrorHandling
export default class FuncNode extends PureComponent<Props, State> {
  public static defaultProps: Partial<Props> = {
    declarationID: '',
  }

  constructor(props) {
    super(props)
    this.state = {
      isExpanded: false,
    }
  }

  public render() {
    const {
      func,
      bodyID,
      service,
      onChangeArg,
      declarationID,
      onGenerateScript,
      declarationsFromBody,
    } = this.props
    const {isExpanded} = this.state

    return (
      <div
        className="func-node"
        onMouseEnter={this.handleMouseEnter}
        onMouseLeave={this.handleMouseLeave}
      >
        <div className="func-node--name">{func.name}</div>
        <FuncArgsPreview func={func} />
        {isExpanded && (
          <FuncArgs
            func={func}
            bodyID={bodyID}
            service={service}
            onChangeArg={onChangeArg}
            declarationID={declarationID}
            onGenerateScript={onGenerateScript}
            onDeleteFunc={this.handleDelete}
            declarationsFromBody={declarationsFromBody}
          />
        )}
      </div>
    )
  }

  private handleDelete = (): void => {
    const {func, bodyID, declarationID} = this.props

    this.props.onDelete({funcID: func.id, bodyID, declarationID})
  }

  private handleMouseEnter = (e: MouseEvent<HTMLElement>): void => {
    e.stopPropagation()

    this.setState({isExpanded: true})
  }

  private handleMouseLeave = (e: MouseEvent<HTMLElement>): void => {
    e.stopPropagation()

    this.setState({isExpanded: false})
  }
}
