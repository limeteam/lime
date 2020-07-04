import request from './request'

export function getDistributorLevel(id) {
  return request({
    url: '/admin/distributorlevels/' + id,
    method: 'GET'
  })
}

export function DistributorLevelList(query) {
  return request({
    url: '/admin/distributorlevels',
    method: 'GET',
    params: query
  })
}

export function updateDistributorLevel(id, data) {
  return request({
    url: '/admin/distributorlevels/' + id,
    method: 'put',
    data: data
  })
}

export function updatestatus(id, data) {
  return request({
    url: '/admin/distributorlevels/' + id + '/status',
    method: 'post',
    data: data
  })
}

export function createDistributorLevel(data) {
  return request({
    url: '/admin/distributorlevels',
    method: 'post',
    data: data
  })
}

export function deleteDistributorLevel(id,data) {
  return request({
    url: '/admin/distributorlevels/' + id,
    method: 'delete',
    data: data
  })
}
