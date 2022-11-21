class AdminController < ApplicationController
  def index
    # Check for access token
    token = cookies[:access_token]
    unless token.nil?
      user_data = AdminClient.new.get_users(token, 0)
      unless user_data.nil?
        # session[:user_id] = user_data['id']
        @users = User.new(user_data)
      else
        redirect_to user_index_path
      end
    else
      redirect_to user_index_path
    end
  end
end
