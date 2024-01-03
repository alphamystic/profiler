# Imports required for the sample - Google Auth and API Client Library Imports.
# Get these packages from https://pypi.org/project/google-api-python-client/ or run $ pip
# install google-api-python-client from your terminal
from google.oauth2 import service_account
from googleapiclient import _auth

SCOPES = ['https://www.googleapis.com/auth/chronicle-backstory']

# The apikeys-demo.json file contains the customer's OAuth 2 credentials.
# SERVICE_ACCOUNT_FILE is the full path to the apikeys-demo.json file
# ToDo: Replace this with the full path to your OAuth2 credentials
SERVICE_ACCOUNT_FILE = '/customer-keys/apikeys-demo.json'

# Create a credential using Google Developer Service Account Credential and Chronicle API
# Scope.
credentials = service_account.Credentials.from_service_account_file(SERVICE_ACCOUNT_FILE, scopes=SCOPES)

# Build an HTTP client to make authorized OAuth requests.
http_client = _auth.authorized_http(credentials)

# <your code continues here>
