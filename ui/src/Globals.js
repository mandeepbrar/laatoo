let native = typeof(document) == 'undefined'
let storage = native?  {}: localStorage
let application = native?  {}: document.InitConfig
let wind = native? {}: window
application.native = native

module.exports = {
  Storage: storage,
  Application: application,
  Window: wind
}


//"http://authnserver.falcon.kronos.com/openam/json/authenticate"
//http://bhavna-khurana-back.falcon.kronos.com/
/*
Dim objHTTP As New WinHttp.WinHttpRequest
    objHTTP.Open "POST", AuthenticationServer.Value, False
    objHTTP.setRequestHeader "X-OpenAM-Username", username.Value
    objHTTP.setRequestHeader "X-OpenAM-Password", Password.Value
    objHTTP.setRequestHeader "Content-Type", "application/json"
    objHTTP.send
    If (objHTTP.Status <> 200) Then
    MsgBox "Login unsuccessful."
    Exit Sub
    End If
    headers = objHTTP.getAllResponseHeaders
    Debug.Print headers
    headerlen = Len(headers)
    'Dim cookiesArray(20) As String
    cookiecount = 1
    pos = 1
    Do While pos < headerlen
        foundpos = InStr(pos, headers, "Set-Cookie:")
        If foundpos = 0 Then
            Exit Do
        End If
        semipos = InStr(foundpos, headers, ";")
        If semipos = 0 Then
            Exit Do
        End If
        cookie = Mid(headers, foundpos + 12, semipos - foundpos - 11)
        If Right$(cookie, 1) = ";" Then
            cookie = Left$(cookie, Len(cookie) - 1)
        End If
        cookiesArray(cookiecount) = cookie
        pos = semipos
        cookiecount = cookiecount + 1
    Loop
    count = cookiecount
    success = True
    Exit Sub
ErrorHandler:
    MsgBox "Login unsuccessful ." + Err.Description
End Sub
Private Sub GetAuditData(ByRef auditData As String, ByRef cookiesArray() As String, cookiecount As Integer, ByRef success As Boolean)
    On Error GoTo ErrorHandler:
    parametersString = ""
    If RecCount.Value <> "" Then
    parametersString = parametersString + "count=" + RecCount.Value + "&"
    End If
    If Actor.Value <> "" Then
    parametersString = parametersString + "username=" + Actor.Value + "&"
    End If
    If AuditType.Value <> "" Then
    parametersString = parametersString + "audittype=" + AuditType.Value + "&"
    End If
    If StartDate.Value <> "" Then
    stDate = CDate(StartDate.Value)
    parametersString = parametersString + "auditBeginDate=" + Format(stDate, "yyyy-mm-dd") + "&"
    End If
    If EndDate.Value <> "" Then
    EndDate = CDate(EndDate.Value)
    parametersString = parametersString + "auditEndDate=" + Format(EndDate, "yyyy-mm-dd") + "&"
    End If
    URL = HostName.Value + "wfc/restcall/audit/v1/audititem/auditrecords?" + parametersString
    Dim auditReq As New WinHttp.WinHttpRequest
    'Dim auditReq As New Microsoft.XMLHTTP
    AUTHN_TOKEN = ""
    JWK_URI = ""
    auditReq.Open "GET", URL, False

    For i = 1 To cookiecount - 1
        auditReq.setRequestHeader "Cookie", cookiesArray(i)
    Next i
    auditReq.setRequestHeader "Content-Type", "application/json"
    auditReq.send
    If (auditReq.Status <> 200) Then
    MsgBox "Could not fetch audit data from server"
    Exit Sub
    End If
    success = True
    auditData = auditReq.responseText
    Exit Sub
ErrorHandler:
    MsgBox "Could not fetch audit data from server." + Err.Description
End Sub*/
