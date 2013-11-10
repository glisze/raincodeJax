## RaincodeJax

TeX Formulae with [MathJax](http://raincodejax.appthing.com/) and ECMAscript
for [Google App Engine](https://appengine.google.com/).

Author: Gunter Liszewski@gmail.com


## Project setup

1. Install the [App Engine Python SDK](https://developers.google.com/appengine/downloads)


## Testing the app locally

To run the app locally:

```
dev_appserver.py .
```


## Deploying

To deploy the application:

1. Use the [Google Cloud Console](https://cloud.google.com/console) to create a project
1. Replace `your-app-id` in `app.yaml` with the project id from the previous step
1. Deploy the application:

```
appcfg.py --oauth2 update .
```


## Contributing changes

See [CONTRIB.md](CONTRIB.md)


# Licensing

See [LICENSE](LICENSE)

   Copyright 2013 Gunter Liszewski

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
