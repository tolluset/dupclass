// Test component with function call patterns
import { cn } from '../utils/cn';
import clsx from 'clsx';

const Component2 = () => {
  return (
    <div>
      {/* cn function - duplicate: rounded-lg */}
      <div className={cn("rounded-lg shadow-md rounded-lg")}>cn function duplicate</div>
      
      {/* clsx function - duplicate: border */}
      <div className={clsx('border border-gray-300 border')}>clsx function duplicate</div>
      
      {/* Multiple duplicates in one line */}
      <div className="mt-4 mb-4 mt-4 mb-4 px-2 py-2">Multiple duplicates</div>
      
      {/* Edge cases */}
      <div className="">Empty className</div>
      <div className="   ">Only spaces</div>
      <div className="single-class">Single class</div>
      
      {/* Should be ignored - not className */}
      <div class="not-className">Wrong attribute</div>
      <div id="some-id">No className attribute</div>
    </div>
  );
};

export default Component2;