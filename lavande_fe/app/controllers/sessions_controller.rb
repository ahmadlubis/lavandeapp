class SessionsController < ApplicationController

  def new
    if logged_in
      redirect_to user_index_path
    end
  end

  def create
    response = UserClient.new.login(session_params) 
    unless response.nil?
      session[:is_logged_in] = true
      # cookies[:access_token] = response['access_token']
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
      flash[:alert] = "An error occured when logging in"
    end  
  end

  def destroy
    session[:is_logged_in] = false
    session[:user_id] = nil
    session[:role] = nil
    cookies[:access_token] = nil
    redirect_to new_session_path
  end
  
  private

  def session_params
    params.require(:email)
    params.require(:password)
    params.permit(:email, :password)
  end
end
