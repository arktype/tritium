require 'nokogiri'
require_relative '../parser/reader'
require_relative '../parser/preprocess'

module Tritium
  module Engines
    class Base
      XML_PARSERS = {"xml" =>  Nokogiri::XML, "html" =>  Nokogiri::HTML, "xhtml" => Nokogiri::XML}
      
      attr :root_instruction

      def initialize(script_string, options = {})
        xml_parser_name = options["xml_parser"] || options[:xml_parser]
        @script_path = options["path"] || options[:path] || File.dirname(__FILE__)
        @script_string = script_string
        @xml_parser = XML_PARSERS[xml_parser_name || "xml"] || throw("Invalid parser")
        @script_name = options[:script_name] || options["script_name"] || "MAIN"
        @root_instruction = reader_klass.new.read(processed_script)
      end
      
      def reader_klass
        Tritium::Parser::Reader
      end
      
      def processed_script
        Tritium::Parser::Preprocess::run(@script_string, @script_path, @script_name)
      end
      
      def to_script
        @root_instruction.to_script
      end
    end
  end
end
