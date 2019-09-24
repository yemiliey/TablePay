import { get } from 'object-path'
import axios from 'axios'

const app = '/app'
const PAGE_PATHS = {
  API_PROXY: '/api',
  ROOT: app,
  OLD_NEW_APP: '/app2',
  APP_BID_SUBPATH: 'biz',
  // _X_ means this is a function
  _APP_BID_: bid => `${app}/biz/${bid}`,
  _APP_BID_TAB_: (bid, tab) => `${app}/biz/${bid}/${tab}`,
}

export const getEnv = () => {
  // prod | sand | stag | local
  return process.env.REACT_APP_ENV || 'local'
}
export const EnvTypes = {
  PROD: 'prod',
  STAG: 'stag',
  SAND: 'sand',
  Local: 'local',
}

export function getHostByEnv(type, env) {
  const hosts = {
    sand: 'localhost:8080',
    stag: 'nwomni.uw2.zs.cntr.io',
    prod: 'nwomni.uw2.zp.cntr.io'
  }

  if (env === 'local') {
    // local testing assumes sand
    env = 'sand'
  }
  const host = get(hosts, `${type}.${env}`)
  return `${PAGE_PATHS.API_PROXY}/${host}`
}

export const timeout = async ms => {
  return new Promise(resolve => setTimeout(resolve, ms))
}

export async function fetchApi({ method, data, url, headers, ...otherParams }) {
  headers = headers || {}
  return axios({
    method,
    data,
    url,
    ...otherParams,
    timeout: 60000,
    headers: {
      ...headers,
    },
  })
    .then(response => {
      // TODO: intercept meta error conditions from server
      if (response.status < 200 || response.status > 399) {
        throw new Error('HTTP ' + response.status + ': ' + response.statusText)
      }
      return response.data
    })
    .catch(error => {
      // Intercepts network errors for logging, then re-throws it
      // as part of normal promise chain

      // TODO - should log this error somewhere
      // only send 4/500s

      // eslint-disable-next-line no-console
      console.log('axios network error', error.response)
      throw error
    })
}

// expected params: url
// optional params: params
export async function getApi(config) {
  return fetchApi(
    Object.assign({}, config, {
      method: 'GET',
    })
  )
}

// expected params: url
// optional params: params
export async function deleteApi(config) {
  return fetchApi(
    Object.assign({}, config, {
      method: 'DELETE',
    })
  )
}

// Use this for normal POSTs
// expected params: url, data
export async function postApi(config) {
  if (config.data == null) {
    throw new Error('Must call PostApi with data payload')
  }

  return fetchApi(
    Object.assign({}, config, {
      method: 'POST',
    })
  )
}

// expected params: url, data
export async function putApi(config) {
  if (config.data == null) {
    throw new Error('Must call PutApi with data payload')
  }
  return fetchApi(
    Object.assign({}, config, {
      method: 'PUT',
    })
  )
}

// Use this for uploading files, images, etc
// currently used for Campaign AdImage upload
// expected params: url, data instanceof FormData
export async function uploadApi(config) {
  if (config.data == null) {
    throw new Error('Must call PostFormDataApi with data payload')
  }
  if (!(config.data instanceof FormData)) {
    throw new Error('PostFormDataApi data payload must be instance of FormData')
  }
  return fetchApi(
    Object.assign({}, config, {
      method: 'POST',
      config: {
        headers: { 'content-type': 'multipart/form-data' },
      },
    })
  )
}
