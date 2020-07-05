import request from './request'

export function getDistributor(id) {
  return request({
    url: '/admin/distributors/' + id,
    method: 'GET'
  })
}

export function distributorList(query) {
  return request({
    url: '/admin/distributors',
    method: 'GET',
    params: query
  })
}

export function updateDistributor(id, data) {
  return request({
    url: '/admin/distributors/' + id,
    method: 'put',
    data: data
  })
}

export function updatestatus(id, data) {
  return request({
    url: '/admin/distributors/' + id + '/status',
    method: 'post',
    data: data
  })
}

export function createDistributor(data) {
  return request({
    url: '/admin/distributors',
    method: 'post',
    data: data
  })
}

export function deleteDistributor(id, data) {
  return request({
    url: '/admin/distributors/' + id,
    method: 'delete',
    data: data
  })
}

export function distributionSetting(code, data) {
  return request({
    url: '/admin/configs/' + code + '/distributor',
    method: 'post',
    data: data
  })
}

export function getDistributionSetting(code, data) {
  return request({
    url: '/admin/distributor/' + code,
    method: 'GET'
  })
}