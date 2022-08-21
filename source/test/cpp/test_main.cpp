#include "cbase/c_target.h"
#include "cbase/c_allocator.h"
#include "xunittest/xunittest.h"

#include "xfibers/x_fibers.h"

UNITTEST_SUITE_LIST(xFibersUnitTest);
UNITTEST_SUITE_DECLARE(xFibersUnitTest, fibers);

namespace ncore
{
	class TestHeapAllocator : public alloc_t
	{
	public:
							TestHeapAllocator(ncore::alloc_t* allocator) : mAllocator(allocator) { 		}

		virtual void*		v_allocate(u32 size, u32 alignment)
		{
			UnitTest::IncNumAllocations();
			return mAllocator->allocate(size, alignment);
		}

		virtual u32			v_deallocate(void* mem)
		{
			UnitTest::DecNumAllocations();
			return mAllocator->deallocate(mem);
		}

		virtual void		v_release()		{ }
	private:
		ncore::alloc_t*	mAllocator;
	};
}

class UnitTestAllocator : public UnitTest::Allocator
{
public:
					UnitTestAllocator(ncore::alloc_t* allocator) : mAllocator(allocator) {}
	virtual void*	Allocate(size_t size)		{ return mAllocator->allocate((ncore::u32)size, 4);}
	virtual size_t	Deallocate(void* ptr)		{ return mAllocator->deallocate(ptr); }
private:
	ncore::alloc_t*	mAllocator;
};

ncore::alloc_t* gSystemAllocator = nullptr;
ncore::alloc_t* gAtomicAllocator = nullptr;

bool gRunUnitTest(UnitTest::TestReporter& reporter)
{
	gSystemAllocator = ncore::alloc_t::get_system();
	UnitTestAllocator unittestAllocator( gSystemAllocator );
	UnitTest::SetAllocator(&unittestAllocator);

	ncore::TestHeapAllocator threadHeapAllocator(gSystemAllocator);
	gAtomicAllocator = &threadHeapAllocator;
	
	ncore::atomic::x_Init(gAtomicAllocator);
	int r = UNITTEST_SUITE_RUN(reporter, xFibersUnitTest);
	ncore::atomic::x_Exit();

	gAtomicAllocator = nullptr;

	UnitTest::SetAllocator(nullptr);
	gSystemAllocator->release();
	gSystemAllocator = nullptr;

	return r==0;
}

