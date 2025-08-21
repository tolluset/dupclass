// Test component with duplicate classes
import React from 'react';

const Component1 = () => {
  return (
    <div>
      {/* Double quotes - duplicate: flex */}
      <div className="flex flex items-center">Double quotes duplicate</div>
      
      {/* Single quotes - duplicate: px-4 */}
      <div className='px-4 py-2 px-4 bg-blue-500'>Single quotes duplicate</div>
      
      {/* Backticks - duplicate: bg-red-500 */}
      <div className={`bg-red-500 text-white bg-red-500`}>Backtick duplicate</div>
      
      {/* Normal case - no duplicates */}
      <div className="w-full h-full border rounded">Normal case</div>
    </div>
  );
};

export default Component1;