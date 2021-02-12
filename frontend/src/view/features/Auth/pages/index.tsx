import React from 'react'
import { Switch, Route } from 'react-router-dom'
import { AuthLayout } from './../layouts/AuthLayout'
import LoginPage from './LoginPage'
import RegistrationPage from './RegistrationPage'
import ResendEmailPage from './ResendEmailPage'
import ForgotPasswordPage from './ForgotPasswordPage'

const AuthPages = () => (
  <Switch>
    <Route path="/resend-confirm-email" exact component={ResendEmailPage} />
    <AuthLayout>
      {/* <Route exact path="/" component={LoginPage} /> */}
      <Route path="/login" exact component={LoginPage} />
      <Route path="/registration" exact component={RegistrationPage} />
      <Route path="/forgot-password" exact component={ForgotPasswordPage} />
    </AuthLayout>
  </Switch>
)

export default AuthPages