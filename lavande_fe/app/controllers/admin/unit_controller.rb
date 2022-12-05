class Admin::UnitController < ApplicationController
  require_relative '../../clients/admin/unit_client'
  PAGINATION_LIMIT = 10

  def index
    unit_input = unit_list_query
    if unit_input[:tower].present? && unit_input[:floor].present? && unit_input[:unit_no].present?
      @units ||= session[:units]
      unit_input[:limit] = 50
      result = Admin::UnitClient.new(@token).index(unit_input)
      p result
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
    payload = params.permit(:gov_id, :tower, :floor, :unit_no)

    session[:unit_data] = payload
    result = Admin::UnitClient.new(@token).create(payload)
    if result.success?
      redirect_to admin_unit_index_path, notice: "Successfully created unit %s" % payload[:gov_id]
      session.delete(:unit_data)
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
    payload = params.require(:unit).permit(:id, :gov_id, :tower, :floor, :unit_no)
    payload[:id] = payload[:id].to_i

    result = Admin::UnitClient.new(@token).update(payload)
    if result.success?
      redirect_to admin_unit_index_path, notice: "Successfully updated unit %s data" % payload['gov_id']
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
end
