class SessionsController < ApplicationController
  def create
    token = UserClient.new.login(params[:email], params[:password]) 
    unless token.nil?
      cookies[:access_token] = token
      session[:is_logged_in] = true
      # TODO: Set access token expiry
      # cookies[:access_token] = { value: @response.parsed_response['access_token'], expires: Time.strptime(@response.parsed_response['expired_at']) }

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
    cookies[:access_token] = nil
    redirect_to new_session_path
  end
  
  private

  def user_params
    params.require(:email)
    params.require(:password)
    params.permit(:email, :password)
  end
end
