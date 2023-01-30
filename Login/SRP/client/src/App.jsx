import srp from "@chihuo/srpclient";
import axios from "axios";
import { Buffer } from "buffer";
import { useState } from "react";
export default function App() {
  const username = "chihuo";
  const params = srp.params["2048"];
  const [registerPassword, setRegisterPassword] = useState("123456");
  const [loginPassword, setLoginPassword] = useState("123456");
  //const [verifier, setVerifier] = useState(null);
  const [message, setMessage] = useState("");
  return (
    <div>
      <p>username is "{username}"</p>
      <p>{message}</p>
      <p>
        <input
          value={registerPassword}
          onChange={(event) => {
            setRegisterPassword(event.target.value);
          }}
        />
        <button
          onClick={async (event) => {
            const salt = srp.genKey();
            const u = Buffer.from(username, "utf-8");
            const p = Buffer.from(registerPassword, "utf8");
            const verifier = srp.computeVerifier(params, salt, u, p);
            setMessage("注册成功请登录");
            try {
              let req = {
                username: username,
                salt: salt.toString("hex"),
                verifier: verifier.toString("hex"),
              };
              console.log("reg req", req);
              let resp = await axios.post("/reg", req);
              console.log("reg response", resp);
            } catch (e) {
              console.log(e);
            }
          }}
        >
          注册
        </button>
      </p>
      <p>
        <input
          value={loginPassword}
          onChange={(event) => {
            setLoginPassword(event.target.value);
          }}
        />
        <button
          onClick={async (event) => {
            let secretClient = srp.genKey();

            let client = new srp.Client(params, secretClient);

            let ephemeralA = client.computeA();

            try {
              var req = {
                username: username,
                ephemeralA: ephemeralA.toString("hex"),
              };
              console.log("login/a req", req);
              let iaResponse = await axios.post("/login/a", req);
              console.log("login/a response", iaResponse.data);
              let salt = Buffer.from(iaResponse.data.salt, "hex");
              let ephemeralB = Buffer.from(iaResponse.data.ephemeralB, "hex");
              const u = Buffer.from(username, "utf-8");
              const p = Buffer.from(loginPassword, "utf8");

              client.setPrivate(params, salt, u, p);
              client.setB(ephemeralB);
              let m1 = client.computeM1();

              req = {
                username: username,
                m1: m1.toString("hex"),
              };
              console.log("login/m req", req);
              let loginMResponse = await axios.post("/login/m", req);
              setMessage("登录成功");
            } catch (e) {
              setMessage("!!!登录失败");
              console.log(e);
            }
          }}
        >
          登录
        </button>
      </p>
    </div>
  );
}
