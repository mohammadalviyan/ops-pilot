"use client";

import { LockOutlined, MailOutlined } from "@ant-design/icons";
import { Alert, Button, Card, Form, Input, Space, Typography } from "antd";
import { useRouter } from "next/navigation";
import { useState } from "react";

import { ApiError } from "../../../lib/api";
import { saveAuthSession } from "../../../lib/auth";
import { login } from "../services/auth-service";
import type { LoginRequest } from "../../../types/auth";

const { Text } = Typography;

export function LoginForm() {
  const router = useRouter();
  const [form] = Form.useForm<LoginRequest>();
  const [isSubmitting, setIsSubmitting] = useState(false);
  const [errorMessage, setErrorMessage] = useState<string | null>(null);

  async function handleFinish(values: LoginRequest) {
    setIsSubmitting(true);
    setErrorMessage(null);

    try {
      const result = await login(values);
      saveAuthSession({
        accessToken: result.access_token,
        tokenType: result.token_type,
        expiresIn: result.expires_in,
        user: result.user,
      });
      router.push("/dashboard");
    } catch (error) {
      if (error instanceof ApiError && error.status === 401) {
        setErrorMessage("Unable to sign in. Check your email and password.");
        return;
      }

      setErrorMessage("Unable to reach OpsPilot authentication. Please try again.");
    } finally {
      setIsSubmitting(false);
    }
  }

  return (
    <Card className="op-login-panel op-glass-strong" variant="borderless">
      <Space direction="vertical" size={18} style={{ width: "100%" }}>
        <div>
          <Text strong className="op-login-kicker">
            Secure workspace access
          </Text>
          <h2 className="op-login-title">Sign in to OpsPilot</h2>
          <p className="op-login-copy">
            Continue to the RPA operations workspace to review tickets, robot
            health, assets, and follow-ups.
          </p>
        </div>

        {errorMessage ? (
          <Alert
            type="error"
            showIcon
            message="Sign-in failed"
            description={errorMessage}
          />
        ) : null}

        <Form
          form={form}
          layout="vertical"
          requiredMark={false}
          onFinish={handleFinish}
          initialValues={{
            email: "support@opspilot.local",
          }}
        >
          <Form.Item
            label="Email"
            name="email"
            rules={[
              { required: true, message: "Enter your OpsPilot email." },
              { type: "email", message: "Enter a valid email address." },
            ]}
          >
            <Input
              prefix={<MailOutlined />}
              placeholder="support@opspilot.local"
              autoComplete="email"
              size="large"
            />
          </Form.Item>

          <Form.Item
            label="Password"
            name="password"
            rules={[{ required: true, message: "Enter your password." }]}
          >
            <Input.Password
              prefix={<LockOutlined />}
              placeholder="Enter password"
              autoComplete="current-password"
              size="large"
            />
          </Form.Item>

          <Button
            type="primary"
            htmlType="submit"
            loading={isSubmitting}
            block
            size="large"
          >
            Sign in
          </Button>
        </Form>

        <Text type="secondary" className="op-login-helper">
          Demo support account: support@opspilot.local
        </Text>
      </Space>
    </Card>
  );
}
