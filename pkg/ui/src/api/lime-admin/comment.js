import request from './request'
import qs from 'qs'

export function CommentList() {
  return request({
    url: 'Comments',
    method: 'GET'
  })
}

export function updateComment(id, data) {
  return request({
    url: 'Comments/' + id,
    method: 'put',
    data: data
  })
}

export function createComment(data) {
  return request({
    url: 'Comments',
    method: 'post',
    data: data
  })
}

export function deleteComment(id,data) {
  return request({
    url: 'Comments/' + id,
    method: 'delete',
    data: data
  })
}