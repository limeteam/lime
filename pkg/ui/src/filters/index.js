// set function parseTime,formatTime to filter
// import store from '@/store'

export {
  parseTime,
  formatTime
}
  from '@/utils'
import {
  accMul,
  accDiv
} from '@/utils' // accAdd, accSub, accMul, accDiv

function pluralize(time, label) {
  if (time === 1) {
    return time + label
  }
  return time + label + 's'
}

export function timeAgo(time) {
  const between = Date.now() / 1000 - Number(time)
  if (between < 3600) {
    return pluralize(~~(between / 60), ' minute')
  } else if (between < 86400) {
    return pluralize(~~(between / 3600), ' hour')
  } else {
    return pluralize(~~(between / 86400), ' day')
  }
}

/* 数字 格式化*/
export function numberFormatter(num, digits) {
  const si = [{
    value: 1E18,
    symbol: 'E'
  },
  {
    value: 1E15,
    symbol: 'P'
  },
  {
    value: 1E12,
    symbol: 'T'
  },
  {
    value: 1E9,
    symbol: 'G'
  },
  {
    value: 1E6,
    symbol: 'M'
  },
  {
    value: 1E3,
    symbol: 'k'
  }
  ]
  for (let i = 0; i < si.length; i++) {
    if (num >= si[i].value) {
      return (num / si[i].value + 0.1).toFixed(digits).replace(/\.0+$|(\.[0-9]*[1-9])0+$/, '$1') + si[i].symbol
    }
  }
  return num.toString()
}

export function toThousandFilter(num) {
  return (+num || 0).toString().replace(/^-?\d+/g, m => m.replace(/(?=(?!\b)(\d{3})+$)/g, ','))
}

// 浮点数 * 100
export function accMul100(num) {
  num = Number(num)
  return accMul(num, 100)
}

// 浮点数 / 100
export function accDiv100(num) {
  num = Number(num)
  return accDiv(num, 100)
}

/**
 * 根据货币名字
 * 获取货币symbol id label value
 */
export function getCurrencyInfoBySymbol(currencyName) {
  // const currentCurrencyId = store.getters.currency.currentCurrencyId
  // 获取 store 里面的汇率对照表
  // const currencyInfo = store.getters.moTo.countryCurrencyParams || []
  // for (let i = currencyInfo.length - 1; i >= 0; i--) {
  //   if (currencyInfo[i].name === currencyName) {
  //     return currencyInfo[i].symbol
  //   }
  // }
  // TODO: 临时修改，等服务全部切换完成再处理初始化数据获取
  return currencyName
}
