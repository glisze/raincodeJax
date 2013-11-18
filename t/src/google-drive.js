/***
     2013-11-15, 

     The ECMAscript code is taken from the Google Drive quickstart
     example. The ECMAscript is separated from the HTML to help with
     a jasmine behaviour specification. The aim is to extend and adapt
     is use, while working within a validating environment.

  */
/**
   Global Data
   CLIENT_ID has been generated with the Google Cloud console
   SCOPES     Specifies the function that are permitted using this
  */
        var CLIENT_ID = '996416180033.apps.googleusercontent.com';
        var SCOPES = 'https://www.googleapis.com/auth/drive';
/**
    Start the authorisation flow
    The handler is assigned to the onload event of the google client api script
    (how to validate this facility?)
   */
      function handleClientLoad() {
        window.setTimeout(checkAuth, 1);
      }
/**
    Verify authorisation.
       */
      function checkAuth() {
        gapi.auth.authorize(
            {'client_id': CLIENT_ID, 'scope': SCOPES, 'immediate': true},
            handleAuthResult);
      }
/**
       A callback from the authorization server
       @param {Object} authResult Authorization result
    */
      function handleAuthResult(authResult) {
        var authButton = document.getElementById('authorizeButton');
        var filePicker = document.getElementById('filePicker');
        authButton.style.display = 'none';
        filePicker.style.display = 'none';
        if (authResult && !authResult.error) {
          // successfully authorised. Ready to request api services
          filePicker.style.display = 'block';
          filePicker.onchange = uploadFile;
        } else {
          // Not authorised. Show the button to allow the user to initiate authorisation
          authButton.style.display = 'block';
          authButton.onclick = function() {
              gapi.auth.authorize(
                  {'client_id': CLIENT_ID, 'scope': SCOPES, 'immediate': false},
                  handleAuthResult);
          };
        }
      }
/**
       File upload.
       @param {Object} evt Arguments from the file selector.
   */
      function uploadFile(evt) {
        gapi.client.load('drive', 'v2', function() {
          var file = evt.target.files[0];
          insertFile(file);
        });
      }
/**
    Insert new file.
    @param {File} fileData           File object to read data from.
    @param {Function} callback  Function to call when the request is complete.
  */
      function insertFile(fileData, callback) {
        const boundary = '-------314159265358979323846';
        const delimiter = "\r\n--" + boundary + "\r\n";
        const close_delim = "\r\n--" + boundary + "--";

        var reader = new FileReader();
        reader.readAsBinaryString(fileData);
        reader.onload = function(e) {
          var contentType = fileData.type || 'application/octet-stream';
          var metadata = {
            'title': fileData.name,
            'mimeType': contentType
          };

          var base64Data = btoa(reader.result);
          var multipartRequestBody =
              delimiter +
              'Content-Type: application/json\r\n\r\n' +
              JSON.stringify(metadata) +
              delimiter +
              'Content-Type: ' + contentType + '\r\n' +
              'Content-Transfer-Encoding: base64\r\n' +
              '\r\n' +
              base64Data +
              close_delim;

          var request = gapi.client.request({
              'path': '/upload/drive/v2/files',
              'method': 'POST',
              'params': {'uploadType': 'multipart'},
              'headers': {
                'Content-Type': 'multipart/mixed; boundary="' + boundary + '"'
              },
              'body': multipartRequestBody});
          if (!callback) {
            callback = function(file) {
              console.log(file)
            };
          }
          request.execute(callback);
        }
      }

