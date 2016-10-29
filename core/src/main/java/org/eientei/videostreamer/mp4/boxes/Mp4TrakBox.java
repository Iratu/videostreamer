package org.eientei.videostreamer.mp4.boxes;

import io.netty.buffer.ByteBuf;
import org.eientei.videostreamer.mp4.Mp4Box;
import org.eientei.videostreamer.mp4.Mp4RemuxerHandler;
import org.eientei.videostreamer.mp4.Mp4Track;

import java.util.List;

/**
 * Created by Alexander Tumin on 2016-10-22
 */
public class Mp4TrakBox extends Mp4Box {
    private final Mp4TkhdBox tkhd;
    private final Mp4MdiaBox mdia;

    public Mp4TrakBox(Mp4RemuxerHandler context, List<Mp4Track> tracks, Mp4Track track) {
        super("trak", context);
        this.tkhd = new Mp4TkhdBox(context, tracks, track);
        this.mdia = new Mp4MdiaBox(context, tracks, track);
    }

    @Override
    protected void doWrite(ByteBuf out) {
        tkhd.write(out);
        mdia.write(out);
    }
}

