import request from './request'

export function getUser(id) {
  return request({
    url: '/users/' + id,
    method: 'GET'
  })
}

export function userList(query) {
  return request({
    url: '/admin/users',
    method: 'GET',
    params: query
  })
}

export function updateUser(id, data) {
  return request({
    url: '/admin/users/' + id,
    method: 'put',
    data: data
  })
}

export function updatestatus(id, data) {
  return request({
    url: '/admin/users/' + id + '/status',
    method: 'post',
    data: data
  })
}

export function createUser(data) {
  return request({
    url: '/admin/users',
    method: 'post',
    data: data
  })
}

export function uploadAvatar(formData) {
  return request({
    url: '/admin/upload/face',
    method: 'post',
    data: formData,
    headers: {
      'Content-Type': 'multipart/form-data'
     }
  })
}
