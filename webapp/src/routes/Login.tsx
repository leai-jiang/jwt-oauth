import * as React from "react";
import { withRouter } from "react-router";
import * as querystring from "querystring";
import request from "../utils/request";

interface LoginProps {
    location: any;
}
const Login = (props: LoginProps):JSX.Element => {
    const { location } = props;
    const { code } = querystring.parse(location.search.replace("?", ""));

    async function getUserInfo(code: string) {
        request({
            action: "",
            payload: {}
        })
    }

    React.useEffect(() => {
        if (code) {
            getUserInfo(code as string);
        }
    }, []);

    return (
        <div className="login-div">
            <ul>
                <li>
                    <a href="/api/login/github">
                        <i className="iconfont github"/>
                        <span>Github</span>
                    </a>
                </li>
            </ul>
        </div>
    )
};

export default withRouter(Login);
