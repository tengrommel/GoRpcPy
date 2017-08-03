# coding:utf-8

import json
import socket
import itertools


class RPCClient(object):
    def __init__(self, addr, codec=json):
        self._socket = socket.create_connection(addr)
        self._id_iter = itertools.count()
        self._codec = codec

    # 建立消息
    def _message(self, name, *params):
        # print self._id_iter
        # print type(self._id_iter)
        return dict(id=self._id_iter.next(), params=list(params), method=name)

    # 调用rpc
    def call(self, name, *params):
        # 请求 name="Json.Getname" *params=mv mv=dict(name="di", age=24)
        req = self._message(name, *params)
        # type dict {'params': [{'age': 24, 'name': 'di'}], 'method': 'Json.Getname', 'id': 0}
        id = req.get('id')

        mesg = self._codec.dumps(req)
        # 字典转化为json {"params": [{"age": 24, "name": "di"}], "method": "Json.Getname", "id": 0}
        self._socket.sendall(mesg)

        # This will actually have to loop if resp is bigger
        resp = self._socket.recv(4096)

        resp = self._codec.loads(resp)

        if resp.get('id') != id:
            raise Exception("expected id=%s, received id=%s: %s"
                            %(id, resp.get('id'), resp.get('error')))

        if resp.get('error') is not None:
            raise Exception(resp.get('error'))
        return resp.get('result')

    def close(self):
        self._socket.close()

if __name__ == '__main__':
    rpc = RPCClient(("192.168.1.123", 9789))
    v = dict(name="红雀", age=34)
    print rpc.call("Json_type.Get", v)
    vv = dict(name="海鹰", age=49)
    print rpc.call("Json_type.Find", vv)
    vvv = dict(name="红雀", number=34, counter=34)
    vvv = rpc.call("JJ.Inc_Number", vvv)
    vvv = rpc.call("JJ.Inc_Counter", vvv)
    print "vvv:", vvv
    print rpc.call("JJ.Find", vvv)