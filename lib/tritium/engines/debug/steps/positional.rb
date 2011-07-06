class Tritium::Engines::Debug
  class Step::Node
    
    def insert_tag(tag_name, attributes = {})
      inject("<#{tag_name} />")
    end
    
    def inject(html)
      execute_children_on(position_node(node, html).first)
    rescue
      puts html.inspect
    end
    
    def insert(*args)
      if args.size > 1 && !args.first.include?("<")
        insert_tag(*args)
      else
        inject(*args)
      end
    end
  end
end