package org.eientei.videostreamer.mp4;

import io.netty.buffer.ByteBuf;

/**
 * Created by Alexander Tumin on 2016-10-22
 */
public class Mp4AudioTrakMp3 extends Mp4Track {
    public Mp4AudioTrakMp3(Mp4Context context, int channels, int sampleRate, int sampleCount, ByteBuf slice) {
        super(context, 1, 0, 0, sampleRate, sampleCount);
    }

    @Override
    public boolean isKnown() {
        return true;
    }

    @Override
    public void update(ByteBuf readonly, boolean b) {

    }

    @Override
    public void release() {

    }

    @Override
    public String getShortHandler() {
        return null;
    }

    @Override
    public String getLongHandler() {
        return null;
    }

    @Override
    public Mp4Box getInit() {
        return null;
    }

    @Override
    public Mp4Box getMhd() {
        return null;
    }
}
