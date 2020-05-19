import request from './request'
import qs from 'qs'

export function CommentList() {
  return request({
    url: '/comics/comments',
    method: 'GET'
  })
}

export function deleteComment(id,data) {
  return request({
    url: '/comics/comments/' + id,
    method: 'delete',
    data: data
  })
}
