import request from '@/utils/request'

export function fetchList(query) {
  return request({
    url: '/posts/list',
    method: 'get',
    params: query
  })
}
