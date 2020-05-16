import request from './request'
import qs from 'qs'

export function CommentList() {
  return request({
    url: 'comments',
    method: 'GET'
  })
}

export function deleteComment(id,data) {
  return request({
    url: 'comments/' + id,
    method: 'delete',
    data: data
  })
}