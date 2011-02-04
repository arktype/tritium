require_relative '../scope'

module Tritium
  module Scope
    class Attribute < Base
      def initialize(attribute, root, parent)
        @attribute = attribute
        super
      end
  
      def remove
        @attribute.remove
      end
  
      def name(&block)
        @attribute.name = open_text_scope_with(@attribute.name, &block)
      end
  
      def value(&block)
        @attribute.value = open_text_scope_with(@attribute.value, &block)
      end
    end
  end
end