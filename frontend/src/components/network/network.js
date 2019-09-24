import { getApi, getEnv, getHostByEnv, postApi, putApi, timeout } from './utils'

const ENV = getEnv()
let HOST = getHostByEnv('omni', ENV) // https://medialize.github.io/URI.js/about-uris.html
let DEBUG = false

export const getOrder = async () => {
    const postMockData = async () => {
      throw new Error('Not sure how to mock this yet')
    }
    const postRealData = async () => {
      try {
        const response = await postApi({
          url: HOST + `/v2/urlHere`,
          data: {},
        })
        return response.audiences
      } catch (err) {
        // eslint-disable-next-line no-console
        console.log(err)
        throw new Error('getOrder : ' + err.message)
      }
    }
    if (DEBUG) {
      return postMockData()
    } else {
      return postRealData()
    }
  }