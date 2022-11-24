class SessionsController < ApplicationController
  skip_before_action :check_session, only: [:new, :create]

  def new
    if !session[:login_data].nil?
      session[:login_data].as_json().each do |name, value|
        params[name] = value
      end
      session.delete(:login_data)
    end
  end

  def create
    # response = UserClient.new.login(session_params) 
    # unless response.nil?
    #   session[:is_logged_in] = true
    #   # cookies[:access_token] = response['access_token']
    #   p response['access_token']
    #   cookies[:access_token] = { value: response['access_token'], expires: Time.parse(response['expired_at']) }

    #   # if params[:remember_name]
    #   #   cookies[:email] = user.email
    #   #   cookies[:password] = user.password
    #   # else
    #   #   cookies.delete(:email)
    #   #   cookies.delete(:password)
    #   # end

    #   redirect_to user_index_path, notice: "Logged in"
    # else
    #   flash[:alert] = "An error occured when logging in"
    # end
    login_data = session_params
    session[:login_data] = login_data.except(:password)
    result = UserClient.new.login(login_data)
    if result.success?
      session[:is_logged_in] = true
      response = result.parsed_response
      p response['access_token']
      cookies[:access_token] = { value: response['access_token'], expires: Time.parse(response['expired_at']) }

      # if params[:remember_name]
      #   cookies[:email] = user.email
      #   cookies[:password] = user.password
      # else
      #   cookies.delete(:email)
      #   cookies.delete(:password)
      # end

      redirect_to user_index_path, notice: "Logged in"
    else
      err_msg = result.parsed_response['error_message']
      redirect_back fallback_location: new_session_path, alert: "Error: %s" % err_msg
    end
  end

  def destroy
    session[:is_logged_in] = false
    session[:user_id] = nil
    session[:role] = nil
    cookies.delete(:access_token)
    redirect_to new_session_path, notice: "Successfully logged out"
  end
  
  private

  def session_params
    params.require([:email, :password])
    params.permit(:email, :password)
  end
end
