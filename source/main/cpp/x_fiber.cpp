namespace xfibers
{
    typedef void (*ffunc_t)(void*);

    enum ePRIORITY
    {
        PRIORITY_HIGH,
        PRIORITY_NORMAL,
        PRIORITY_LOW,
    };

    enum eIOCMD
    {
        IO_OPEN,
        IO_CLOSE,
        IO_READ,
        IO_OPENREAD,
        IO_OPENREADCLOSE,
    };

    struct iodata_t
    {
        byte*       m_buffer;
        u32         m_size;
    };

    struct fiber_t
    {
        // fiber context
        u32   m_findex;
        u32   m_fstack_size;
        void* m_fstack;
        void* m_fcontext;

        // fiber local params (can change when a fiber is resumed)
        volatile params_t* m_params;
        
        volatile byte* m_iodata;

        // conditional
        counter_t* m_counter;
        fiber_t*   m_resume;
    };

    struct iocmd_t
    {
        fiber_t*  m_fiber_to_resume;
        ePRIORITY m_fiber_priority;

        const char* m_name;
        byte*       m_read_buffer;
        byte const* m_write_buffer;
        eIOCMD      m_iocmd;
        u32         m_offset;
        u32         m_size;
    };

    fiber_t* create(size_t stack_size, ffunc_t f, void* arg);
    void     destroy(fiber_t* fiber);
    void     switch_to(fiber_t* fiber);
    bool     convert_to_thread(void);
    fiber_t* convert_from_thread(void* arg);

    void run_async(fiber_t* fiber);
    void run_sync(fiber_t* fiber);
    void run_all_async(fiber_t** fibers, u32 count);
    void run_all_sync(fiber_t** fibers, u32 count);

    void read_file_sync(const char* filename, byte* buffer, u32 offset, u32 size, xfibers::ePRIORITY priority);
    void write_file_sync(const char* filename, byte const* buffer, u32 offset, u32 size, xfibers::ePRIORITY priority);

    void read_file_async(const char* filename, byte* buffer, u32 offset, u32 size, xfibers::ePRIORITY priority, xfibers::func_t fiber_func);
    void write_file_async(const char* filename, byte const* buffer, u32 offset, u32 size, xfibers::ePRIORITY priority, xfibers::func_t fiber_func);
} // namespace xfibers

#define FIBER_PROC(fiber_func) fiber_func(void* arg)

FIBER_PROC(process)
{
    // xfibers::get_params();

    // This call will post the IO request and switch to another fiber, when the IO request
    // has finished (or failed), the processing will come back here.
    byte* map_buffer = allocate_buffer(1024*1024*8);
    read_file_sync("main.map", map_buffer, 0, 1024*1024*8, xfibers::PRIORITY_HIGH);

    // Process the mapdata
}

FIBER_IO_PROC(process_mapdata)
{
    // xfibers::get_params();
    // xfibers::get_iodata();
}

FIBER_PROC(load_map)
{
    // xfibers::get_params();

    // This call will post the IO request and switch to another fiber, when the IO request
    // has finished (or failed), the processing will come back here.
    // When passing nullptr for the buffer it will be allocated by the IO thread.
    read_file_async("main.map", nullptr, 0, 0, xfibers::PRIORITY_HIGH, process_mapdata);
}
