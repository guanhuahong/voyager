const LOC_REGX = /^(?:http|https):\/\/(?:127\.0\.0\.1|localhost)/
const LAN_REGX = /^(?:http|https):\/\/(?:192\.168(?:\.[0-9]{1,3}){2})/

export const isLocalhost = LOC_REGX.test(location.href)

export const isLan = LAN_REGX.test(location.href) || isLocalhost