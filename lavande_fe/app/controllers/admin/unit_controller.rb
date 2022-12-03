class Admin::UnitController < ApplicationController
  require_relative '../../clients/admin/unit_client'
  PAGINATION_LIMIT = 10

  def index
    unit_input = unit_list_query
    if unit_input[:tower].present? && unit_input[:floor].present? && unit_input[:unit_no].present?
      @units ||= session[:units]
      result = Admin::UnitClient.new(@token).index(unit_input)
      if result.success?
        unless result.parsed_response['data'].empty?
          @cur_unit = result.parsed_response['data'][0]
        end
      else
        err_msg = result.parsed_response['error_message']
        redirect_back fallback_location: admin_unit_index_path, alert: "An error occurred when fetching units: %s" % err_msg
      end
    elsif unit_input[:tower].present? && unit_input[:floor].present?
      @units ||= []
      result = Admin::UnitClient.new(@token).index(unit_input)
      if result.success?
        for unit in result.parsed_response['data'] do
          @units << [unit['gov_id'], unit['unit_no'].to_i]
        end
        session[:units] = @units
      else
        err_msg = result.parsed_response['error_message']
        redirect_back fallback_location: admin_unit_index_path, alert: "An error occurred when fetching units: %s" % err_msg
      end
    end
  end

  def new
    unless session[:unit_data].nil?
      session[:unit_data].as_json().each do |name, value|
        params[name] = value
      end
      session.delete(:unit_data)
    end
  end

  def create
    unit_data = create_unit_payload
    session[:unit_data] = unit_data
    result = Admin::UnitClient.new(@token).create(unit_data)
    if result.success?
      redirect_to admin_unit_index_path, notice: "Successfully created unit %s" % unit_data[:gov_id]
    else
      err_msg = result.parsed_response['error_message']
      redirect_back fallback_location: admin_unit_index_path, alert: "An error occurred when creating unit: %s" % err_msg
    end
  end

  def edit
    result = Admin::UnitClient.new(@token).index(edit_unit_query)
    if result.success? && !result.parsed_response['data'].empty?
      unit_data = result.parsed_response['data'][0]
      @unit = Unit.new(unit_data)
    else
      redirect_to admin_unit_index_path, alert: "An error occurred when retrieving unit data"
    end
  end

  def update
    unit_data = update_unit_payload
    result = Admin::UnitClient.new(@token).update(unit_data)
    if result.success?
      redirect_to admin_unit_index_path, notice: "Successfully updated unit %s data" % unit_data['gov_id']
    else
      err_msg = result.parsed_response['error_message']
      redirect_back fallback_location: admin_unit_index_path, alert: "An error occurred when updating unit data: %s" % err_msg
    end
  end

  private

  def unit_list_query
    query = {}
    if params[:tower].present?
      query[:tower] = params[:tower]
    end
    if params[:floor].present?
      query[:floor] = params[:floor]
    end
    if params[:unit_no].present?
      query[:unit_no] = params[:unit_no]
    end
    query[:page] = 1
    query[:limit] = PAGINATION_LIMIT
    query[:offset] = 0
    query
  end

  def create_unit_payload
    params.require([:gov_id, :tower, :floor, :unit_no])
    payload = params.permit(:gov_id, :tower, :floor, :unit_no)
    payload
  end

  def edit_unit_query
    params.require(:id)
    query = params.permit(:id)
    # query[:gov_id] = query[:id]
    # query.delete(:id)
    query[:page] = 1
    query[:limit] = PAGINATION_LIMIT
    query[:offset] = 0
    query
  end

  def update_unit_payload
    payload = params.require(:unit).permit(:id, :gov_id, :tower, :floor, :unit_no)
    payload[:id] = payload[:id].to_i
    payload
  end
end
