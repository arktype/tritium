

module Tritium
  module Engines
    class Standard < Base
      
      def initialize(script_string, options = {})
        super
        compile!
      end
      
      def compile!
        
      end
    end
  end
end