class AdminController < ApplicationController
  # before_action :check_token

  def index
    # token = cookies[:access_token]
    # unless token.nil?
    #   user_data = AdminClient.new.get_users(token, params[:search], 0)
    #   unless user_data.nil?
    #     # session[:user_id] = user_data['id']
    #     @users ||= []
    #     for user in user_data do
    #       @users << User.new(user)
    #     end
    #   else
    #     redirect_to user_index_path
    #   end
    # else
    #   redirect_to new_session_path
    # end
    user_data = AdminClient.new.get_users(@token, params[:search], 0)
    unless user_data.nil?
      # session[:user_id] = user_data['id']
      @users ||= []
      for user in user_data do
        @users << User.new(user)
      end
    else
      redirect_to user_index_path
    end
  end

  def status
    # unless token.nil?
    #   result = AdminClient.new.change_status(@token, params[:id], params[:action])
    #   unless result.nil?
    #     flash[:notice] = "Successfully updated user status"
    #     # redirect_to admin_index_path
    #     redirect_back(fallback_location: admin_index_path)
    #   else
    #     flash[:alert] = "An error occurred when updating user status"
    #     redirect_back(fallback_location: admin_index_path)
    #   end
    # else
    #   flash[:alert] = "An error occurred when updating user status"
    #   # redirect_to admin_index_path
    #   redirect_back(fallback_location: admin_index_path)
    # end
    status_data = status_params
    status_data[:target_id] = status_data[:target_id].to_i
    result = AdminClient.new.change_status(@token, status_data)
    # unless result.nil?
    #   flash[:notice] = "Successfully updated user status"
    #   # redirect_to admin_index_path
    #   redirect_back(fallback_location: admin_index_path)
    # else
    #   flash[:alert] = "An error occurred when updating user status"
    #   redirect_back(fallback_location: admin_index_path)
    # end
    if result.success?
      redirect_back fallback_location: admin_index_path, notice: "Successfully updated user status of %s" % params[:name]
    else
      err_msg = result.parsed_response['error_message']
      redirect_back fallback_location: admin_index_path, alert: "An error occurred when updating user status: %s" % err_msg
    end
  end

  private

  def status_params
    params.require([:target_id, :status])
    params.permit(:target_id, :status)
  end
end
