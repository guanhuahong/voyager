/**
 * 遮罩身份证号码
 */
export const coverIDCode = makeMarkFilter(6, 14, '*')
/**
 * 遮罩手机号码
 */
export const coverPhoneNum = makeMarkFilter(3, 7,'*')

export const columnFormatWithIdno = (column, row, cellValue, index) => {
    return coverIDCode(cellValue)
}

export const columnFormatWithPhone = (column, row, cellValue, index) => {
    return coverPhoneNum(cellValue)
}
/**
 * 
 * @param {number} start 开始序号
 * @param {number} end 技术序号 
 * @param {string} separator 替换字符 
 * @returns Fillter 函数
 */
export function makeMarkFilter(start, end, separator) {
    if (end < start) {
        [start, end] = [end, start]
    }
    if (!separator) separator = '*'
    return (value) => {
        if (typeof value === "number") value = value.toString()
        if (typeof value !== 'string') return value
        
        var strLen = value.length,
            sepLen = end - start
        if (strLen < start) return value
        var result = ''
        if (start > 0) {
            var ss = 0,
                se = start
            result = value.slice(ss, se)
        }
        if (start + sepLen < strLen) {
            result += separator.repeat(sepLen)
        } else {
            strLen > start && (result += separator.repeat(strLen - start))
            return result
        }
        
        if (end > start && end < strLen) {
            result += value.slice(end, strLen)
        }
        return result
    }
}