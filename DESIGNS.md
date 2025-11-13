## Creating a New Logbook

### Using the **LogbookApp**

#### Overview

1. User creates a new logbook. This is created in the local database.
2. Once created, the user has the option to register the logbook with the Station Manager online service.
   - If the logbook is NOT registered with the Station Manager online service, none of the QSO logged under that
   logbook will be forwarded (or able to be forwarded/uploaded) to the Station Manager online service.
   - Each logbook has its own unique API key associated with it.
   - The API key is used to authenticate the logbook with the Station Manager online service.
3. The user can continue to use the logbook locally to log QSOs under that logbook, and QSO will be
uploaded/forwarded to other configured online services (e.g. QRZ.com)

#### Logbook Registration

1. User clicks on the **Register** button (at the time of logbook creation) or later (when an Internet connection is available)..
2. The app requests a new API key from the Station Manager online service.
3. The app then registers the logbook with the Station Manager online service and associates the API key with the logbook.
4. The API key is saved locally in the local database associated with that logbook.
5. All QSOs logged under that logbook will be uploaded/forwarded to the Station Manager online service.
6. The user can now continue to use the logbook locally to log QSOs under that logbook, and QSO will be
   uploaded/forwarded to the Station-Manager online services as well as other configured online services (e.g. QRZ.com)
