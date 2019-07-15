#### 添加权限验证

本地测试
```
pipenv run python run.py
```

测试命令  
`headers`必须带上`Authorization: Bearer (jwtoken)`
```
curl http://127.0.0.1:5000/channel/list -H 'Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOiIxOSIsInVuYW1lIjoid2VuZGVsbCIsImV4cCI6MTU1NDAyNTcyOH0.XcU63XBQVp6RFUes4gh2xqc7FaNAskiv64L6AvfGSu6pc_tI4WJ_jnp3Q5Vl-8RuV4ugdubiNg01Bk-l-IoBDCz7PBpJH-6-ezcQAWkV7P-4adtD9bnSOMVID_yfVnAcsyyNTamG43ZavQ2MM2F3G0oDq_Iz--vMv5c3clAtQRM'
```

#### 上线前需要与权限中心对接