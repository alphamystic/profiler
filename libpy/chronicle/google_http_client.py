# connector.py

"""
This package exposes a connector that returns a google HTTP client for  chronicle clients:
    It takes in a string pointing to the location for the json file for credentials and returns an error or a
"""


from google.oauth2 import service_account
from googleapiclient import _auth

class GoogleHTTPClient:
    SCOPES = ['https://www.googleapis.com/auth/chronicle-backstory']
    @classmethod
    def get_http_client(cls, service_account_file):
        try:
            credentials = service_account.Credentials.from_service_account_file(service_account_file, scopes=cls.SCOPES)
            http_client = _auth.authorized_http(credentials)
            return http_client, None
        except Exception as e:
            return None, f"Error: {e}"

# Example usage:
"""
from connector import  GoogleHTTPClient as ghc
SERVICE_ACCOUNT_FILE = '/customer-keys/apikeys-demo.json'
http_client, error = ghc.get_http_client(SERVICE_ACCOUNT_FILE)

if error:
    print("Error occurred:", error)
else:
    print("HTTP client:", http_client)
"""
