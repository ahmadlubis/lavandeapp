class Admin::TenantController < ApplicationController
  PAGINATION_LIMIT = 10

  def index
    @result = {}
    
    # User
    user_client = Admin::UserClient.new(@token)
    if params[:user_id].present?
      user = user_client.index(selected_user_query)
      if user.success?
        @result["user"] = user.parsed_response["data"][0]
      else
        err_msg = user.parsed_response['error_message']
        redirect_back fallback_location: admin_tenant_index_path, alert: "An error occurred when fetching selected user: %s" % err_msg
      end
    end
    
    if params[:name].present? || params[:email].present?
      users = user_client.index(user_list_query)
      if users.success?
        @result["users"] = users.parsed_response
      else
        err_msg = users.parsed_response['error_message']
        path = admin_tenant_index_path
        if params[:id]
          path = admin_tenant_index_path(params[:id])
        end
        redirect_back fallback_location: path, alert: "An error occurred when fetching users: %s" % err_msg
      end
    end
    
    # Unit
    unit_client = Admin::UnitClient.new(@token)
    if params[:unit_id].present?
      unit = unit_client.index(selected_unit_query)
      if unit.success?
        @result["unit"] = unit.parsed_response["data"][0]
      else
        err_msg = unit.parsed_response['error_message']
        redirect_back fallback_location: admin_tenant_index_path, alert: "An error occurred when fetching selected unit: %s" % err_msg
      end
    end
    
    if params[:tower].present? || params[:floor].present?
      units = unit_client.index(unit_list_query)
      if units.success?
        @result["units"] = units.parsed_response
      else
        err_msg = units.parsed_response['error_message']
        path = admin_tenant_index_path
        if params[:id]
          path = admin_tenant_index_path(params[:id])
        end
        redirect_back fallback_location: path, alert: "An error occurred when fetching units: %s" % err_msg
      end
    end

    # Tenants
    if params[:user_id].present? && params[:unit_id].present?
      tenants = Admin::TenantClient.new(@token).index(tenant_list_query)
      if tenants.success?
        @result["tenants"] = tenants.parsed_response
      else
        err_msg = tenants.parsed_response['error_message']
        redirect_back fallback_location: unit_index_path, alert: "An error occurred when fetching tenants: %s" % err_msg
      end

      @result["tenants"]["data"].each do |v|
        user = user_client.index({id: v["user_id"], limit: 1, offset: 0})
        if user.success?
          v["user"] = user.parsed_response["data"][0]
        else
          err_msg = tenants.parsed_response['error_message']
          redirect_back fallback_location: unit_index_path, alert: "An error occurred when fetching tenants: %s" % err_msg
        end

        unit = unit_client.index({id: v["unit_id"], limit: 1, offset: 0})
        if unit.success?
          v["unit"] = unit.parsed_response["data"][0]
        else
          err_msg = tenants.parsed_response['error_message']
          redirect_back fallback_location: unit_index_path, alert: "An error occurred when fetching tenants: %s" % err_msg
        end
      end
    end
  end

  private

  def selected_user_query
    query = params.permit(:user_id)
    query[:id] = query.delete :user_id
    query[:limit] = 1
    query[:offset] = 0
    query
  end

  def user_list_query
    params[:user_page] ||= 1
    params[:user_page] = params[:user_page].to_i
    query = params.permit(:user_page, :name, :email)
    query[:page] = query.delete :user_page
    query[:limit] = PAGINATION_LIMIT
    query[:offset] = (query[:page] - 1) * PAGINATION_LIMIT
    query
  end

  def selected_unit_query
    query = params.permit(:unit_id)
    query[:id] = query.delete :unit_id
    query[:limit] = 1
    query[:offset] = 0
    query
  end

  def unit_list_query
    params[:unit_page] ||= 1
    params[:unit_page] = params[:unit_page].to_i
    query = params.permit(:unit_page, :tower, :floor)
    query[:page] = query.delete :unit_page
    query[:limit] = PAGINATION_LIMIT
    query[:offset] = (query[:page] - 1) * PAGINATION_LIMIT
    query
  end

  def tenant_list_query
    params[:tenant_page] ||= 1
    params[:tenant_page] = params[:tenant_page].to_i
    query = params.permit(:tenant_page, :unit_id, :user_id)
    query[:page] = query.delete :tenant_page
    query[:limit] = PAGINATION_LIMIT
    query[:offset] = (query[:page] - 1) * PAGINATION_LIMIT
    query
  end
end