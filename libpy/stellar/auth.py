import requests
import base64
import json
from urllib.parse import urlunparse
requests.packages.urllib3.disable_warnings()

# Step 1
# Add DP IP/hostname, userid, and refresh token from GUI here
HOST = "cyflare.stellarcyber.cloud"
userid = "myuser@stellarcyber.ai"
refresh_token = "2iRpBAyQYEfv77R2QtATlJN6Nvq6uzftBdzotSy2pjT-IvJTLw9aiHyh7Y2mo12IDSWc-FfHwUyPpmiHQnJrSH"

def getAccessToken(userid, refresh_token):
    auth = base64.b64encode(bytes(userid + ":" + refresh_token, "utf-8")).decode("utf-8")
    headers = {
        "Authorization": "Basic " + auth,
        "Content-Type": "application/x-www-form-urlencoded",
    }
    url = urlunparse(("https", HOST, "/connect/api/v1/access_token", "", "", ""))
    res = requests.post(url, headers=headers, verify=False)
    print(res.status_code)
    return res.json()["access_token"]


def getIncidents(token):
    headers = {"Authorization": "Bearer " + token}
    url = urlunparse(("https", HOST, "/connect/api/v1/incidents?tenant_id=817418af76d44358922636e34be9627c?sort=incident_score&order=desc&limit=10", "", "", ""))
    res = requests.get(url, headers=headers, verify=False)
    print(res.status_code)
    return res.json()

if __name__ == "__main__":

    # Step 2: Use getAccessToken with supplied credentials to generate JWT
    jwt = getAccessToken(userid, refresh_token)
    print("------------ jwt -------------")
    print(jwt)
    print("------------ jwt  end -------------")

    # Step 3: use JWT token to call public API
    incidents = getIncidents(jwt)
    print("------------ call result of /connect/api/v1/incidents -------------")
    print(incidents)
    print("------------ end api results -------------")
