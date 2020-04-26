import { DocumentContext } from 'next/document'
import { Store } from 'redux'
import { RootState } from '~/modules/rootState'

/**
 * NextDocumentContext with redux store context
 * @tree
 */
export type AppContext = DocumentContext & {
  readonly store: Store<RootState>
  readonly auth: {
    token?: string
  }
}
