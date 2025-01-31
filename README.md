## simple ldc api in golang
- [x] operators work
- [x] operators also used for dates (FormDate, DateVisited)
- [x] POST request handling with operators
- [x] limit/offset work
- [x] aws cognito group discrimination works
- [x] containerized 
- [x] CICD-ready
- [ ] unified logs with levels across the app
- [x] exceptions to the routes (like tblProject)
- [x] add aero data handling

### upticks
v1.0.6
- allowed enum on `tblProject`

v1.0.5
- enabled cors middleware on main

v1.0.4 
- geoindicators static schema at `/internal/schemas/geoIndicators.go` still referenced `"Precipitation_Yearly_Maximum_Daily_2_YR"`

v1.0.3 
- geoindicators static schema at `/internal/schemas/geoIndicators.go` still referenced `wkb_geometry`

v1.0.2
- removed columns and columntypes db request

v1.0.1
- exception to middleware routing
- aerodata handling


### to use, include config.yaml at root level with this structure:
```yaml
server:
  port: 8080

database:
  host: x.x.x.x
  port: xxxx
  name: dbname
  tenants:
    public:
      user: user1
      password: xxx
    publication:
      user: user2
      password: xxx
    legal:
      user: user3
      password: xxx

awsCognito:
  userPoolId: x
  clientId: x
  tokenType: id 
```