/* @flow */
import engine from '../engine'
import {log} from './log/logui'
import type {IncomingCallMapType} from '../constants/types/flow-types'

export default function ListenLogUi () {
  engine.listenOnConnect('ListenLogUi', () => {
    const params: IncomingCallMapType = {
      'keybase.1.logUi.log': (params, response) => {
        log(params)
        response.result()
      }
    }

    engine.listenGeneralIncomingRpc(params)
    console.log('Registered Listener for logUi.log')
  })
}
