package handlers

const headerContentType = "Content-Type"
const headerAccessControlAllowOrigin = "Access-Control-Allow-Origin"
const headerAccessControlAllowMethods = "Access-Control-Allow-Methods"
const headerAccessControlAllowHeader = "Access-Control-Allow-Headers"
const headerAccessControlExposeHeader = "Access-Control-Expose-Headers"
const headerXFrameOption = "X-Frame-Options"
const headerXForwarded = "X-Forwarded-For"

const contentTypeJSON = "application/json"
const contentTypeText = "text/plain"
const contentTypeHTML = "text/html"

// Twilio token and account SID
// TODO: Add to a .gitignorefile
const trialNum = "+14252121598"
const serviceSID = "VAf592ae55ca6059b22f2f6255d668bd98"
const twilAuthString = "https://verify.twilio.com/v2/Services/" + serviceSID + "/Verifications"
const dcMsg = `You are receiving this message because New-Era has lost contact with your device. If this was not planned, please contact us immediatly.`
