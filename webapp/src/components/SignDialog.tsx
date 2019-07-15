import * as React from "react";
import * as ReactDOM from "react-dom";
import { Modal, Icon } from "antd";
import "./SignDialog.css";

interface SignDialogProps {
  visible: boolean;
  title?: string;
  onCancel?: any;
  className?: string;
}

function signInByGithub() {
  window.location.href = "http://localhost:5005/api/login/github";
}
const SignDialog = (props: SignDialogProps) => {
  return (
    <Modal
      {...props}
      footer={null}
      className="sign-dialog"
      width={300}
    >
      <div className="sign-dialog-container">
        <Icon type="github" onClick={() => signInByGithub()} />
        <Icon type="wechat" className="wechat-icon" />
      </div>
    </Modal>
  )
};

function signIn() {
  const div = document.createElement("div");

  function render() {
    document.body.appendChild(div);
    ReactDOM.render(<SignDialog visible title="Sign in" onCancel={destroy} />, div)
  }

  function destroy() {
      ReactDOM.unmountComponentAtNode(div);
      if (div) {
        document.body.removeChild(div);
      }
  }

  render();
}

export { SignDialogProps, SignDialog, signIn as default };
